package igdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCreditTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	cr := Credit{}
	typ := reflect.ValueOf(cr).Type()

	err := c.validateStruct(typ, CreditEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCredit(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		ID     int
		ExpErr string
	}{
		{"Happy path", "test_data/get_credit.txt", 1342182279, ""},
		{"Invalid ID", "test_data/empty.txt", -321, ErrNegativeID.Error()},
		{"Empty Response", "test_data/empty.txt", 1342182279, errEndOfJSON.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.GetCredit(tt.ID)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			eID := 1342182279
			aID := cr.ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Michael"
			an := cr.Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := CreditCategory(5)
			ac := cr.Category
			if ac != ec {
				t.Errorf("Expected category %d, got %d", ec, ac)
			}

			ep := 45
			ap := cr.Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}

func TestGetCredits(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		IDs    []int
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/get_credits.txt", []int{1342181334, 1342186852}, []OptionFunc{OptLimit(5)}, ""},
		{"Invalid ID", "test_data/empty.txt", []int{-100}, nil, ErrNegativeID.Error()},
		{"Zero IDs", "test_data/empty.txt", nil, nil, ErrEmptyIDs.Error()},
		{"Empty Response", "test_data/empty.txt", []int{1342181334, 1342186852}, nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", []int{1342181334, 1342186852}, []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.GetCredits(tt.IDs, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 2
			al := len(cr)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1342180316
			aID := cr[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Scott"
			an := cr[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := CreditCategory(5)
			ac := cr[1].Category
			if ac != ec {
				t.Errorf("Expected category %d, got %d", ec, ac)
			}

			ep := 140
			ap := cr[1].Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}

func TestSearchCredits(t *testing.T) {
	var creditTests = []struct {
		Name   string
		Resp   string
		Qry    string
		Opts   []OptionFunc
		ExpErr string
	}{
		{"Happy path", "test_data/search_credits.txt", "jim", []OptionFunc{OptLimit(50)}, ""},
		{"Empty query", "test_data/search_credits.txt", "", []OptionFunc{OptLimit(50)}, ""},
		{"Empty response", "test_data/empty.txt", "jim", nil, errEndOfJSON.Error()},
		{"Invalid option", "test_data/empty.txt", "jim", []OptionFunc{OptOffset(9999)}, ErrOutOfRange.Error()},
	}
	for _, tt := range creditTests {
		t.Run(tt.Name, func(t *testing.T) {
			ts, c, err := testServerFile(http.StatusOK, tt.Resp)
			if err != nil {
				t.Fatal(err)
			}
			defer ts.Close()

			cr, err := c.SearchCredits(tt.Qry, tt.Opts...)
			assertError(t, err, tt.ExpErr)

			if tt.ExpErr != "" {
				return
			}

			el := 3
			al := len(cr)
			if al != el {
				t.Errorf("Expected length of %d, got %d", el, al)
			}

			eID := 1342181334
			aID := cr[0].ID
			if aID != eID {
				t.Errorf("Expected ID %d, got %d", eID, aID)
			}

			en := "Justin - Mom Cody Mark Josh Jim Kerri"
			an := cr[0].Name
			if an != en {
				t.Errorf("Expected name '%s', got '%s'", en, an)
			}

			ec := 1463521290038
			ac := cr[1].CreatedAt
			if ac != ec {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", ec, ac)
			}

			eu := 1463521290038
			au := cr[1].UpdatedAt
			if au != eu {
				t.Errorf("Expected Unix time in milliseconds of %d, got %d", eu, au)
			}

			eCat := CreditCategory(5)
			aCat := cr[2].Category
			if aCat != eCat {
				t.Errorf("Expected category %d, got %d", eCat, aCat)
			}

			ep := 365
			ap := cr[2].Position
			if ap != ep {
				t.Errorf("Expected position %d, got %d", ep, ap)
			}
		})
	}
}
