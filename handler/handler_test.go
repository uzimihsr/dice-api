package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/uzimihsr/dice-api/mock_dice"
)

// リクエストパスが"/"でクエリパラメータがない場合, faces=6でRollが呼び出されることのテスト
func TestDiceGet01(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(1)
	m.EXPECT().Roll().Return(3).Times(1)
	m.EXPECT().GetFaces().Return(6).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/", DiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 3
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/"で"faces=12"をクエリパラメータで指定した場合, faces=12でRollが呼び出されることのテスト
func TestDiceGet02(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(12).Times(1)
	m.EXPECT().Roll().Return(3).Times(1)
	m.EXPECT().GetFaces().Return(12).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/", DiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?faces=12", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 3
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/"で"faces=-12"をクエリパラメータで指定した場合, faces=6でRollが呼び出されることのテスト
func TestDiceGet03(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(1)
	m.EXPECT().Roll().Return(3).Times(1)
	m.EXPECT().GetFaces().Return(6).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/", DiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?faces=-12", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 3
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/"で"faces=asdf"をクエリパラメータで指定した場合, エラーになることのテスト
func TestDiceGet04(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(0)
	m.EXPECT().Roll().Return(3).Times(0)
	m.EXPECT().GetFaces().Return(6).Times(0)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/", DiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?faces=asdf", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusInternalServerError {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// リクエストパスが"/cheat"でクエリパラメータがない場合, num=6でCheatが呼び出されることのテスト
func TestCheatDiceGet01(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(1)
	m.EXPECT().Cheat(6).Return(6).Times(1)
	m.EXPECT().GetFaces().Return(6).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 6
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/cheat"で"number=4"をクエリパラメータで指定した場合, num=4でCheatが呼び出されることのテスト
func TestCheatDiceGet02(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(1)
	m.EXPECT().Cheat(4).Return(4).Times(1)
	m.EXPECT().GetFaces().Return(4).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat?number=4", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 4
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/cheat"で"faces=12"をクエリパラメータで指定した場合, num=12でCheatが呼び出されることのテスト
func TestCheatDiceGet03(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(12).Times(1)
	m.EXPECT().Cheat(12).Return(12).Times(1)
	m.EXPECT().GetFaces().Return(12).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat?faces=12", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 12
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/cheat"で"faces=-12"をクエリパラメータで指定した場合, num=6でCheatが呼び出されることのテスト
func TestCheatDiceGet04(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(1)
	m.EXPECT().Cheat(6).Return(6).Times(1)
	m.EXPECT().GetFaces().Return(6).Times(1)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat?faces=-12", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v", writer.Code)
	}
	result := diceResult{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	expect := 6
	actual := result.Number
	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// リクエストパスが"/cheat"で"faces=asdf"をクエリパラメータで指定した場合, エラーになることのテスト
func TestCheatDiceGet05(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(0)
	m.EXPECT().Cheat(6).Return(6).Times(0)
	m.EXPECT().GetFaces().Return(6).Times(0)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat?faces=asdf", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusInternalServerError {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// リクエストパスが"/cheat"で"number=asdf"をクエリパラメータで指定した場合, エラーになることのテスト
func TestCheatDiceGet06(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(0)
	m.EXPECT().Cheat(6).Return(6).Times(0)
	m.EXPECT().GetFaces().Return(6).Times(0)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/cheat?number=asdf", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusInternalServerError {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// DiceHandlerが無効なメソッドで呼ばれた場合, diceNotFoundが呼ばれることのテスト
func TestDiceNotFound01(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(0)
	m.EXPECT().Roll().Return(3).Times(0)
	m.EXPECT().GetFaces().Return(6).Times(0)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/", DiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusNotFound {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// CheatDiceHandlerが無効なメソッドで呼ばれた場合, diceNotFoundが呼ばれることのテスト
func TestDiceNotFound02(t *testing.T) {

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_dice.NewMockDiceInterface(ctrl)
	m.EXPECT().SetFaces(6).Times(0)
	m.EXPECT().Cheat(6).Return(6).Times(0)
	m.EXPECT().GetFaces().Return(6).Times(0)

	// do
	mux := http.NewServeMux()
	mux.HandleFunc("/cheat", CheatDiceHandler(m))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/cheat", nil)
	mux.ServeHTTP(writer, request)

	// check
	if writer.Code != http.StatusNotFound {
		t.Errorf("Response code is %v", writer.Code)
	}
}
