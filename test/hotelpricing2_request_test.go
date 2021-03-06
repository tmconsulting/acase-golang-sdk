package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelPricingRequest2_Ok(t *testing.T) {
	testRequest("hotelpricing_response_example.xml", false)
	defer gockOff()

	data, err := acApi.HotelPricingRequest2(context.Background(), 0, 0, 0, 0, 0,
		0, 0, 0, 0, "", "", "",
		"", "", "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Code, 1100577)
	st.Expect(t, data.Name, "Авеню Парк Отель")
	st.Expect(t, data.StartDate, "20.10.2011")
	st.Expect(t, data.EndDate, "25.10.2011")
	st.Expect(t, data.NumberOfGuests, 2)
	st.Expect(t, data.NumberOfExtraBedsAdult, 1)
	st.Expect(t, data.NumberOfExtraBedsChild, 0)
	st.Expect(t, data.NumberOfExtraBedsInfant, 0)
	st.Expect(t, data.SpecialRequirements, "1. Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно..")
	st.Expect(t, data.Country.Code, 9)
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Country.Position.Latitude, "57.5158")
	st.Expect(t, data.Country.Position.Longitude, "82.6172")
	st.Expect(t, data.AdmUnit1.Code, 1)
	st.Expect(t, data.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.AdmUnit2.Code, 1)
	st.Expect(t, data.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TypeOfPlace.Code, 1)
	st.Expect(t, data.TypeOfPlace.Name, "город")
	st.Expect(t, data.City.Code, 2)
	st.Expect(t, data.City.Name, "Москва")
	st.Expect(t, data.City.Position.Latitude, "55.758959")
	st.Expect(t, data.City.Position.Longitude, "37.616272")
	st.Expect(t, data.ObjType.Code, 2)
	st.Expect(t, data.ObjType.Name, "Гостиница")
	st.Expect(t, data.Position.Latitude, "55.781619")
	st.Expect(t, data.Position.Longitude, "37.750605")
	st.Expect(t, data.Currency.Code, 2)
	st.Expect(t, data.Currency.Name, "RUB")
	st.Expect(t, data.Infants.AgeFrom, 0)
	st.Expect(t, data.Infants.AgeTo, 2)
	st.Expect(t, data.Infants.UseThisAge.Code, 1)
	st.Expect(t, data.Infants.UseThisAge.Name, "Да")
	st.Expect(t, data.Children.UseThisAge.Code, 2)
	st.Expect(t, data.Children.UseThisAge.Name, "Нет")
	st.Expect(t, len(data.SpecialRequirementList.SpecialRequirement), 1)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Code, 595)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Text, "Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно.")
	st.Expect(t, data.HotelPricingDetail.Code, "1120501")
	st.Expect(t, data.HotelPricingDetail.RoomName, "Стандарт 2-местный (2 кровати)")
	st.Expect(t, data.HotelPricingDetail.RoomCode, "900250")
	st.Expect(t, data.HotelPricingDetail.RateCode, "900092")
	st.Expect(t, data.HotelPricingDetail.RateName, "FIT")
	st.Expect(t, data.HotelPricingDetail.RateDescription, "")
	st.Expect(t, data.HotelPricingDetail.RateGroupCode, "4")
	st.Expect(t, data.HotelPricingDetail.RateGroupName, "FIT")
	st.Expect(t, data.HotelPricingDetail.IsHourRate, "2")
	st.Expect(t, data.HotelPricingDetail.MaxHours, 0)
	st.Expect(t, data.HotelPricingDetail.StartTime, "")
	st.Expect(t, data.HotelPricingDetail.EndTime, "")
	st.Expect(t, data.HotelPricingDetail.MinimumRetailPrice, 12000.00)
	st.Expect(t, data.HotelPricingDetail.Price, 8280.00)
	st.Expect(t, data.HotelPricingDetail.TravelAgencyCommission, 0.00)
	st.Expect(t, data.HotelPricingDetail.FullVATInPrice, "1")
	st.Expect(t, data.HotelPricingDetail.VATIncludedInPrice, 1263.05)
	st.Expect(t, data.HotelPricingDetail.DeadlineDate, "19.10.2011")
	st.Expect(t, data.HotelPricingDetail.DeadlineTimeLoc, "19.10.2011 13:00")
	st.Expect(t, data.HotelPricingDetail.DeadlineTimeSys, "19.10.2011 13:00")
	st.Expect(t, data.HotelPricingDetail.DeadlineTimeUTC, "19.10.2011 10:00")
	st.Expect(t, data.HotelPricingDetail.PenaltySize, 2760.00)
	st.Expect(t, data.HotelPricingDetail.DefaultCheckInTime, "14:00")
	st.Expect(t, data.HotelPricingDetail.DefaultCheckOutTime, "12:00")
	st.Expect(t, data.HotelPricingDetail.CheckInTime, "14:00")
	st.Expect(t, data.HotelPricingDetail.CheckOutTime, "12:00")
	st.Expect(t, data.HotelPricingDetail.MaxNumberOfGuests, 2)
	st.Expect(t, data.HotelPricingDetail.MaxNumberOfExtraBeds, 1)
	st.Expect(t, data.HotelPricingDetail.MaxNumberOfExtraBedsAdult, 1)
	st.Expect(t, data.HotelPricingDetail.MaxNumberOfExtraBedsChild, 0)
	st.Expect(t, data.HotelPricingDetail.MaxNumberOfExtraBedsInfant, 0)
	st.Expect(t, len(data.HotelPricingDetail.SpecialOfferList.Items), 1)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].StayNights, 5)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].PayNights, 4)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].SpecialOfferType.Code, 17)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].SpecialOfferType.Name, "Бесплатные ночи")
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].SpecialOfferType.Id, 3)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].IsMultiple.Code, 2)
	st.Expect(t, data.HotelPricingDetail.SpecialOfferList.Items[0].IsMultiple.Name, "Нет")
	st.Expect(t, data.HotelPricingDetail.AllowEarlierCheckIn.Code, 1)
	st.Expect(t, data.HotelPricingDetail.AllowEarlierCheckIn.Name, "Да")
	st.Expect(t, data.HotelPricingDetail.AllowLateCheckOut.Code, 1)
	st.Expect(t, data.HotelPricingDetail.AllowLateCheckOut.Name, "Да")
	st.Expect(t, data.HotelPricingDetail.AllowToAmendBookings.Code, 1)
	st.Expect(t, data.HotelPricingDetail.AllowToAmendBookings.Name, "Да")
	st.Expect(t, data.HotelPricingDetail.AllowToAmendGuestNames.Code, 1)
	st.Expect(t, data.HotelPricingDetail.AllowToAmendGuestNames.Name, "Да")
	st.Expect(t, data.HotelPricingDetail.AllGuestNamesRequired.Code, 2)
	st.Expect(t, data.HotelPricingDetail.AllGuestNamesRequired.Name, "Нет")
	st.Expect(t, data.HotelPricingDetail.GuestCitizenshipRequired.Code, 2)
	st.Expect(t, data.HotelPricingDetail.GuestCitizenshipRequired.Name, "Нет")
	st.Expect(t, len(data.HotelPricingDetail.Meals.Items), 1)
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].Code, 1101889)
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].Name, "Завтрак \"Континентальный\"")
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].Price, 180.00)
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].MinimumRetailPrice, 350.00)
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].IncludedInPrice.Code, 2)
	st.Expect(t, data.HotelPricingDetail.Meals.Items[0].IncludedInPrice.Name, "Нет")
	st.Expect(t, len(data.HotelPricingDetail.PenaltyRuleList.Items), 1)
	st.Expect(t, data.HotelPricingDetail.PenaltyRuleList.Items[0].HoursFrom, 0)
	st.Expect(t, data.HotelPricingDetail.PenaltyRuleList.Items[0].HoursTo, 24)
	st.Expect(t, data.HotelPricingDetail.PenaltyRuleList.Items[0].Value, 100.0)
	st.Expect(t, data.HotelPricingDetail.PenaltyRuleList.Items[0].CalculationRuleCode, 2)
	st.Expect(t, data.HotelPricingDetail.PenaltyRuleList.Items[0].CalculationRuleName, "% от стоимости размещ. за 1 ночь")
	st.Expect(t, len(data.HotelPricingDetail.WhereToPayList.Items), 1)
	st.Expect(t, data.HotelPricingDetail.WhereToPayList.Items[0].Code, 3)
	st.Expect(t, data.HotelPricingDetail.WhereToPayList.Items[0].Name, "Оплата ООО \"АКАДЕМСЕРВИС\" согласно договору")
	st.Expect(t, data.HotelPricingDetail.PriceDescription.EarlierCheckInPrice.Price, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.EarlierCheckInPrice.TravelAgencyCommission, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.LateCheckOutPrice.Price, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.LateCheckOutPrice.TravelAgencyCommission, 0.00)
	st.Expect(t, len(data.HotelPricingDetail.PriceDescription.Days.Items), 5)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].Date, "20.10.2011")
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].MinimumRetailPrice, 3000.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].Price, 2070.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].TravelAgencyCommission, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].AdditionalMeal, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].ExtraBedsAdultPrice, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].ExtraBedsChildPrice, 0.00)
	st.Expect(t, data.HotelPricingDetail.PriceDescription.Days.Items[0].ExtraBedsInfantPrice, 0.00)
	st.Expect(t, len(data.EarlyCheckInList.Items), 7)
	st.Expect(t, data.EarlyCheckInList.Items[0].Time, "10:00")
	st.Expect(t, data.EarlyCheckInList.Items[0].Price, 3450.00)
	st.Expect(t, data.EarlyCheckInList.Items[0].Text, "")
	st.Expect(t, data.EarlyCheckInList.Items[0].Guaranteed.Code, 1)
	st.Expect(t, data.EarlyCheckInList.Items[0].Guaranteed.Name, "Да")
	st.Expect(t, data.EarlyCheckInList.Items[0].EarlyCheckInLateCheckOutRule.Code, 2)
	st.Expect(t, data.EarlyCheckInList.Items[0].EarlyCheckInLateCheckOutRule.Name, "Доплата")
	st.Expect(t, len(data.LateCheckOutList.Items), 4)
	st.Expect(t, data.LateCheckOutList.Items[0].Time, "12:30 - 14:00")
	st.Expect(t, data.LateCheckOutList.Items[0].Price, 0.00)
	st.Expect(t, data.LateCheckOutList.Items[0].Text, "Заранее не бронируется. Бесплатно при наличии мест в гостинице")
	st.Expect(t, data.LateCheckOutList.Items[0].Guaranteed.Code, 2)
	st.Expect(t, data.LateCheckOutList.Items[0].Guaranteed.Name, "Нет")
	st.Expect(t, data.LateCheckOutList.Items[0].EarlyCheckInLateCheckOutRule.Code, 3)
	st.Expect(t, data.LateCheckOutList.Items[0].EarlyCheckInLateCheckOutRule.Name, "Не определено")
}

func TestHotelPricingRequest_Error(t *testing.T) {
	testRequest("hotelpricing_error_example.xml", true)
	defer gockOff()

	_, err := acApi.HotelPricingRequest2(context.Background(), 0, 0, 0, 0, 0,
		0, 0, 0, 0, "", "", "",
		"", "", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
