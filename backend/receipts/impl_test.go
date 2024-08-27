package receipts

import (
	"encoding/json"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReceipt_datetime(t *testing.T) {
	type fields = Receipt
	tests := []struct {
		name    string
		fields  fields
		want    time.Time
		wantErr bool
	}{
		{
			name: "happy",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "1.29",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			want: time.Date(2022, 7, 13, 17, 26, 0, 0, time.UTC),
		},
		{
			name: "time parse failure seconds",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "1.29",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26:13",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			want:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: true,
		},
		{
			name: "time parse failure hours",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "1.29",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "99:26",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			want:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: true,
		},
		{
			name: "date parse failure",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "1.29",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 14, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26:13",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			want:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Receipt{
				Items:        tt.fields.Items,
				PurchaseDate: tt.fields.PurchaseDate,
				PurchaseTime: tt.fields.PurchaseTime,
				Retailer:     tt.fields.Retailer,
				Total:        tt.fields.Total,
			}
			got, err := r.datetime()
			if (err != nil) != tt.wantErr {
				t.Errorf("datetime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("datetime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceipt_score(t *testing.T) {
	tests := []struct {
		name    string
		fields  Receipt
		want    int
		wantErr bool
	}{
		{
			name: "morning",
			fields: Receipt{
				Retailer:     "Walgreens",
				PurchaseDate: types.Date{Time: time.Date(2022, 01, 02, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "08:13",
				Total:        "2.65",
				Items: []Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
					{ShortDescription: "Dasani", Price: "1.40"},
				},
			},
			want:    15,
			wantErr: false,
		},
		{
			name: "simple",
			fields: Receipt{
				Retailer:     "Target",
				PurchaseDate: types.Date{Time: time.Date(2022, 01, 02, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "13:13",
				Total:        "1.25",
				Items: []Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
				},
			},
			want:    31,
			wantErr: false,
		},
		{
			name: "complex",
			fields: Receipt{
				Retailer:     "Target",
				PurchaseDate: types.Date{Time: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "13:01",
				Total:        "35.35",
				Items: []Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
			},
			want:    28,
			wantErr: false,
		},
		{
			name: "gatorade",
			fields: Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: types.Date{Time: time.Date(2022, 03, 02, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "14:33",
				Total:        "9.00",
				Items: []Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
			},
			want:    109,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Receipt{
				Items:        tt.fields.Items,
				PurchaseDate: tt.fields.PurchaseDate,
				PurchaseTime: tt.fields.PurchaseTime,
				Retailer:     tt.fields.Retailer,
				Total:        tt.fields.Total,
			}
			got, err := r.score()
			if (err != nil) != tt.wantErr {
				t.Errorf("score() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("score() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceipt_validate(t *testing.T) {
	tests := []struct {
		name    string
		fields  Receipt
		wantErr bool
	}{
		{
			name: "valid",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "1.29",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			wantErr: false,
		},
		{
			name: "invalid price letters",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "abcd",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			wantErr: true,
		},
		{
			name: "invalid price mixed",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34a",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "19.37",
			},
			wantErr: true,
		},
		{
			name: "invalid total letters",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "19.37",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "abcd",
			},
			wantErr: true,
		},
		{
			name: "invalid total mixed",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26",
				Retailer:     "Wal*Mart",
				Total:        "19a.37",
			},
			wantErr: true,
		},
		{
			name: "invalid time seconds",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:26:00",
				Retailer:     "Wal*Mart",
				Total:        "19a.37",
			},
			wantErr: true,
		},
		{
			name: "invalid time letters",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "17:2x6",
				Retailer:     "Wal*Mart",
				Total:        "19a.37",
			},
			wantErr: true,
		},
		{
			name: "invalid time hours",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "37:23",
				Retailer:     "Wal*Mart",
				Total:        "19a.37",
			},
			wantErr: true,
		},
		{
			name: "invalid time minutes",
			fields: Receipt{
				Items: []Item{
					{
						ShortDescription: "Great Value Apple Sauce",
						Price:            "12.34",
					},
				},
				PurchaseDate: types.Date{Time: time.Date(2022, 7, 13, 0, 0, 0, 0, time.UTC)},
				PurchaseTime: "37:63",
				Retailer:     "Wal*Mart",
				Total:        "19a.37",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Receipt{
				Items:        tt.fields.Items,
				PurchaseDate: tt.fields.PurchaseDate,
				PurchaseTime: tt.fields.PurchaseTime,
				Retailer:     tt.fields.Retailer,
				Total:        tt.fields.Total,
			}
			if err := r.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetReceiptsIdPoints(t *testing.T) {
	type fields struct {
		Points map[string]int
	}
	type args struct {
		w  *httptest.ResponseRecorder
		r  *http.Request
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				Points: map[string]int{
					"a": 38,
				},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/receipts/a/points", nil),
				id: "a",
			},
			want: 38,
		},
		{
			name: "missing id",
			fields: fields{
				Points: map[string]int{
					"a": 38,
				},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  httptest.NewRequest("GET", "/receipts/b/points", nil),
				id: "b",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Points: tt.fields.Points,
			}
			s.GetReceiptsIdPoints(tt.args.w, tt.args.r, tt.args.id)
			var resp GetReceiptsIdPointsResponse
			err := json.Unmarshal(tt.args.w.Body.Bytes(), &resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReceiptsIdPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := resp.Points
			if got != tt.want {
				t.Errorf("GetReceiptsIdPoints() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErr && tt.args.w.Code != http.StatusOK {
				t.Errorf("GetReceiptsIdPoints() got = %v, want %v", tt.args.w.Code, http.StatusOK)
			} else if tt.wantErr && tt.args.w.Code != http.StatusNotFound {
				t.Errorf("GetReceiptsIdPoints() got = %v, want %v", tt.args.w.Code, http.StatusNotFound)
			}
		})
	}
}

func TestServer_PostReceiptsProcess(t *testing.T) {
	type fields struct {
		Points map[string]int
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				Points: map[string]int{
					"a": 38,
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/receipts/process", strings.NewReader(`{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}`)),
			},
			want:    15,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Points: tt.fields.Points,
			}
			s.PostReceiptsProcess(tt.args.w, tt.args.r)
			var resp PostReceiptsProcessResponse
			err := json.Unmarshal(tt.args.w.Body.Bytes(), &resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostReceiptsProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, ok := s.Points[resp.Id]
			if ok == tt.wantErr {
				t.Errorf("PostReceiptsProcess() ok = %v, wantErr %v", ok, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("PostReceiptsProcess() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErr && tt.args.w.Code != http.StatusCreated {
				t.Errorf("PostReceiptsProcess() got = %v, want %v", tt.args.w.Code, http.StatusCreated)
			} else if tt.wantErr && tt.args.w.Code != http.StatusBadRequest {
				t.Errorf("PostReceiptsProcess() got = %v, want %v", tt.args.w.Code, http.StatusBadRequest)
			}
		})
	}
}
