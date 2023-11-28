package meetingGroupProposals

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
)

/*
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals
{
    public class MeetingGroupProposalStatus : ValueObject
    {
        private MeetingGroupProposalStatus(string value)
        {
            Value = value;
        }

        public static MeetingGroupProposalStatus ToVerify => new MeetingGroupProposalStatus("ToVerify");

        public static MeetingGroupProposalStatus Verified => new MeetingGroupProposalStatus("Verified");

        public string Value { get; }

        internal static MeetingGroupProposalStatus Create(string value)
        {
            return new MeetingGroupProposalStatus(value);
        }
    }
}
*/

// MeetingGroupProposalStatus represents the status of a meeting group proposal.
type MeetingGroupProposalStatus struct {
	domain.ValueObject
	string
}

var (
	// ToVerify indicates that the meeting group proposal needs to be verified.
	meetingGroupProposalStatusToVerify MeetingGroupProposalStatus = MeetingGroupProposalStatus{string: "ToVerify"}
	// Verified indicates that the meeting group proposal has been verified.
	meetingGroupProposalStatusVerified MeetingGroupProposalStatus = MeetingGroupProposalStatus{string: "Verified"}
	// Verified indicates that the meeting group proposal has been rejected.
	meetingGroupProposalStatusRejected MeetingGroupProposalStatus = MeetingGroupProposalStatus{string: "Rejected"}
)

// NewMeetingGroupProposalStatus creates a new MeetingGroupProposalStatus instance.
func NewMeetingGroupProposalStatus(value string) MeetingGroupProposalStatus {
	return MeetingGroupProposalStatus{string: value}
}

func MeetingGroupProposalStatusToVerify() MeetingGroupProposalStatus {
	return meetingGroupProposalStatusToVerify
}

func MeetingGroupProposalStatusVerified() MeetingGroupProposalStatus {
	return meetingGroupProposalStatusVerified
}

func MeetingGroupProposalStatusRejected() MeetingGroupProposalStatus {
	return meetingGroupProposalStatusRejected
}
