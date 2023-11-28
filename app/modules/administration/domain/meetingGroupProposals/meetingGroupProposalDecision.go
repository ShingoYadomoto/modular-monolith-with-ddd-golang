package meetingGroupProposals

import (
	"time"

	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/modules/administration/domain/users"
)

/*
using System;
using CompanyName.MyMeetings.BuildingBlocks.Domain;
using CompanyName.MyMeetings.Modules.Administration.Domain.Users;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals
{
    public class MeetingGroupProposalDecision : ValueObject
    {
        private MeetingGroupProposalDecision(DateTime? date, UserId userId, string code, string rejectReason)
        {
            this.Date = date;
            this.UserId = userId;
            this.Code = code;
            this.RejectReason = rejectReason;
        }

        public DateTime? Date { get; }

        public UserId UserId { get; }

        public string Code { get; }

        public string RejectReason { get; }

        internal static MeetingGroupProposalDecision NoDecision =>
            new MeetingGroupProposalDecision(null, null, null, null);

        private bool IsAccepted => this.Code == "Accept";

        private bool IsRejected => this.Code == "Reject";

        internal static MeetingGroupProposalDecision AcceptDecision(DateTime date, UserId userId)
        {
            return new MeetingGroupProposalDecision(date, userId, "Accept", null);
        }

        internal static MeetingGroupProposalDecision RejectDecision(DateTime date, UserId userId, string rejectReason)
        {
            return new MeetingGroupProposalDecision(date, userId, "Reject", rejectReason);
        }

        internal MeetingGroupProposalStatus GetStatusForDecision()
        {
            if (this.IsAccepted)
            {
                return MeetingGroupProposalStatus.Verified;
            }

            if (this.IsRejected)
            {
                return MeetingGroupProposalStatus.Create("Rejected");
            }

            return MeetingGroupProposalStatus.ToVerify;
        }
    }
}
*/

// MeetingGroupProposalDecision represents the decision for a meeting group proposal.
type MeetingGroupProposalDecision struct {
	domain.ValueObject
	Date         time.Time
	UserID       users.UserID
	Code         string
	RejectReason string
}

// NoDecision represents a default decision with no specific action.
var MeetingGroupProposalDecisionNoDecision = MeetingGroupProposalDecision{}

// NewMeetingGroupProposalDecision creates a new MeetingGroupProposalDecision instance.
func NewMeetingGroupProposalDecision(date time.Time, userID users.UserID, code, rejectReason string) MeetingGroupProposalDecision {
	return MeetingGroupProposalDecision{
		Date:         date,
		UserID:       userID,
		Code:         code,
		RejectReason: rejectReason,
	}
}

// AcceptDecision creates a new decision indicating acceptance.
func AcceptDecision(date time.Time, userID users.UserID) MeetingGroupProposalDecision {
	return NewMeetingGroupProposalDecision(date, userID, "Accept", "")
}

// RejectDecision creates a new decision indicating rejection.
func RejectDecision(date time.Time, userID users.UserID, rejectReason string) MeetingGroupProposalDecision {
	return NewMeetingGroupProposalDecision(date, userID, "Reject", rejectReason)
}

// GetStatusForDecision returns the meeting group proposal status based on the decision.
func (d MeetingGroupProposalDecision) GetStatusForDecision() MeetingGroupProposalStatus {
	if d.IsAccepted() {
		return MeetingGroupProposalStatusVerified()
	}

	if d.IsRejected() {
		return MeetingGroupProposalStatusRejected()
	}

	return MeetingGroupProposalStatusToVerify()
}

// IsAccepted returns true if the decision indicates acceptance.
func (d MeetingGroupProposalDecision) IsAccepted() bool {
	return d.Code == "Accept"
}

// IsRejected returns true if the decision indicates rejection.
func (d MeetingGroupProposalDecision) IsRejected() bool {
	return d.Code == "Reject"
}
