package receipts

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

// Values for alphanumeric regex and Decimals are used in scoring
var alphanumeric = regexp.MustCompile("[[:alnum:]]")
var four, _ = decimal.NewFromFloat64(4)
var oneFifth, _ = decimal.NewFromFloat64(0.2)

// Constants for 2PM and 4PM are used in scoring
const twoPmMinutes = (2 + 12) * 60
const fourPmMinutes = (4 + 12) * 60

type Server struct {
	Points map[string]int
}

type PostReceiptsProcessResponse struct {
	Id string `json:"id"`
}

func (s Server) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// validate the receipt values
	err = receipt.validate()

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// use a type 1 uuid to ensure we don't have to worry about collisions
	id, err := uuid.NewUUID()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// calculate the score
	score, err := receipt.score()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// stash the store
	s.Points[id.String()] = score

	resp := PostReceiptsProcessResponse{Id: id.String()}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (r Receipt) validate() error {
	// Ensure at least one item
	if len(r.Items) == 0 {
		return errors.New("no items found")
	}

	// Ensure the total is a decimal value
	_, err := decimal.Parse(r.Total)
	if err != nil {
		return err
	}

	// Ensure that for each item the price is a decimal value
	for _, i := range r.Items {
		_, err := decimal.Parse(i.Price)
		if err != nil {
			return err
		}
	}

	// Ensure that the purchase date plus purchase time add up to a datetime
	_, err = time.Parse(time.DateTime, r.PurchaseDate.Format(time.DateOnly)+" "+r.PurchaseTime+":00")
	if err != nil {
		return err
	}

	return nil
}

// datetime creates a time object from the receipt date and time
func (r Receipt) datetime() (time.Time, error) {
	return time.Parse(time.DateTime, r.PurchaseDate.Format(time.DateOnly)+" "+r.PurchaseTime+":00")
}

// score calculates the score for a receipt
// returns an error if unable to parse dollar values, unable to manipulate dollar values, or unable to parse datetime
func (r Receipt) score() (int, error) {
	score := 0

	// One point for every alphanumeric character in the retailer name.
	matches := alphanumeric.FindAllStringIndex(r.Retailer, -1)
	if matches != nil {
		score += len(matches)
	}

	// dollar amounts
	total, err := decimal.Parse(r.Total)
	if err != nil {
		return 0, err
	}

	// 50 points if the total is a round dollar amount with no cents.
	if total.IsInt() {
		score += 50
	}

	// 25 points if the total is a multiple of 0.25.
	// calculate total * 4 then determine whether that is an integer
	totalx4, err := total.Mul(four)

	if err != nil {
		return 0, err
	}

	if totalx4.IsInt() {
		score += 25
	}

	// 5 points for every two items on the receipt.
	// order of operations is important
	score += 5 * (len(r.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the
	// nearest integer. The result is the number of points earned.
	for _, item := range r.Items {
		// Short description is a multiple of 3
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			// get the price
			price, err := decimal.Parse(item.Price)
			if err != nil {
				return 0, err
			}

			// multiply by 0.2
			price, err = price.Mul(oneFifth)
			if err != nil {
				return 0, err
			}

			// get the ceiling
			ceil := price.Ceil(0)
			whole, _, ok := ceil.Int64(0)
			if !ok {
				return 0, errors.New("could not retrieve whole part of ceiling value")
			}

			score += int(whole)
		}
	}

	// datetime values
	datetime, err := r.datetime()
	if err != nil {
		return 0, err
	}

	// 6 points if the day in the purchase date is odd.
	if datetime.Day()%2 == 1 {
		score += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	minuteOfDay := datetime.Hour()*60 + datetime.Minute()
	if twoPmMinutes < minuteOfDay && minuteOfDay < fourPmMinutes {
		score += 10
	}

	return score, nil
}

type GetReceiptsIdPointsResponse struct {
	Points int `json:"points"`
}

func (s Server) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	points, ok := s.Points[id]

	if !ok {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}

	resp := GetReceiptsIdPointsResponse{Points: points}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func NewServer() Server {
	return Server{make(map[string]int)}
}
