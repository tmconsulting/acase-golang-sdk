package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestCityListRequest_Ok(t *testing.T) {
	testRequest("citylist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CityListRequest(context.Background(), "", "", 9, 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Countries), 1)
	st.Expect(t, data.Countries[0].Code, 9)
	st.Expect(t, data.Countries[0].Name, "Россия")
	st.Expect(t, data.Countries[0].Position.Latitude, "57.5158")
	st.Expect(t, data.Countries[0].Position.Longitude, "82.6172")
	st.Expect(t, len(*data.Countries[0].Cities), 2)
	st.Expect(t, (*data.Countries[0].Cities)[0].Code, 151)
	st.Expect(t, (*data.Countries[0].Cities)[0].Name, "Абакан")
	st.Expect(t, (*data.Countries[0].Cities)[0].Genitive, "Абакана")
	st.Expect(t, (*data.Countries[0].Cities)[0].Dative, "Абакану")
	st.Expect(t, (*data.Countries[0].Cities)[0].Accusative, "Абакан")
	st.Expect(t, (*data.Countries[0].Cities)[0].Ablative, "Абаканом")
	st.Expect(t, (*data.Countries[0].Cities)[0].Prepositive, "Абакане")
	st.Expect(t, (*data.Countries[0].Cities)[0].AdmUnit1.Code, 1)
	st.Expect(t, (*data.Countries[0].Cities)[0].AdmUnit1.Name, "Не определено")
	st.Expect(t, (*data.Countries[0].Cities)[0].AdmUnit2.Code, 1)
	st.Expect(t, (*data.Countries[0].Cities)[0].AdmUnit2.Name, "Не определено")
	st.Expect(t, (*data.Countries[0].Cities)[0].TypeOfPlace.Code, 2)
	st.Expect(t, (*data.Countries[0].Cities)[0].TypeOfPlace.Name, "город")
	st.Expect(t, (*data.Countries[0].Cities)[0].Position.Latitude, "53.730192")
	st.Expect(t, (*data.Countries[0].Cities)[0].Position.Longitude, "91.454773")
	st.Expect(t, (*data.Countries[0].Cities)[1].Code, 109)
	st.Expect(t, (*data.Countries[0].Cities)[1].Name, "Светлогорск")
	st.Expect(t, (*data.Countries[0].Cities)[1].Genitive, "Светлогорска")
	st.Expect(t, (*data.Countries[0].Cities)[1].Dative, "Светлогорску")
	st.Expect(t, (*data.Countries[0].Cities)[1].Accusative, "Светлогорск")
	st.Expect(t, (*data.Countries[0].Cities)[1].Ablative, "Светлогорском")
	st.Expect(t, (*data.Countries[0].Cities)[1].Prepositive, "Светлогорске")
	st.Expect(t, (*data.Countries[0].Cities)[1].AdmUnit1.Code, 23)
	st.Expect(t, (*data.Countries[0].Cities)[1].AdmUnit1.Name, "Калининградская область")
	st.Expect(t, (*data.Countries[0].Cities)[1].AdmUnit2.Code, 67)
	st.Expect(t, (*data.Countries[0].Cities)[1].AdmUnit2.Name, "Гурьевский район")
	st.Expect(t, (*data.Countries[0].Cities)[1].TypeOfPlace.Code, 2)
	st.Expect(t, (*data.Countries[0].Cities)[1].TypeOfPlace.Name, "город")
	st.Expect(t, (*data.Countries[0].Cities)[1].Position.Latitude, "54.938292")
	st.Expect(t, (*data.Countries[0].Cities)[1].Position.Longitude, "20.158213")
}

func TestCityListRequest_Error(t *testing.T) {
	testRequest("citylist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.CityListRequest(context.Background(), "", "", 9, 0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
