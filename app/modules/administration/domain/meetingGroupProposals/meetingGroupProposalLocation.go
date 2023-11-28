package meetingGroupProposals

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
)

/*
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals
{
    public class MeetingGroupLocation : ValueObject
    {
        private MeetingGroupLocation(string city, string countryCode)
        {
            City = city;
            CountryCode = countryCode;
        }

        public string City { get; }

        public string CountryCode { get; }

        public static MeetingGroupLocation Create(string city, string countryCode)
        {
            return new MeetingGroupLocation(city, countryCode);
        }
    }
}
*/

type MeetingGroupLocation struct {
	domain.ValueObject
	city        string
	countryCode string
}

func NewMeetingGroupLocation(city, countryCode string) MeetingGroupLocation {
	return MeetingGroupLocation{
		city:        city,
		countryCode: countryCode,
	}
}
