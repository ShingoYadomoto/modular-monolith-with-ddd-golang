package meetingGroupProposals

import (
	"time"

	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/modules/administration/domain/users"
)

/*
using System;
using CompanyName.MyMeetings.BuildingBlocks.Domain;
using CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals.Events;
using CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals.Rules;
using CompanyName.MyMeetings.Modules.Administration.Domain.Users;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals
{
    public class MeetingGroupProposal : Entity, IAggregateRoot
    {
        private string _name;

        private string _description;

        private MeetingGroupLocation _location;

        private DateTime _proposalDate;

        private UserId _proposalUserId;

        private MeetingGroupProposalStatus _status;

        private MeetingGroupProposalDecision _decision;

        private MeetingGroupProposal(
            MeetingGroupProposalId id,
            string name,
            string description,
            MeetingGroupLocation location,
            UserId proposalUserId,
            DateTime proposalDate)
        {
            Id = id;
            _name = name;
            _description = description;
            _location = location;
            _proposalUserId = proposalUserId;
            _proposalDate = proposalDate;

            _status = MeetingGroupProposalStatus.ToVerify;
            _decision = MeetingGroupProposalDecision.NoDecision;

            this.AddDomainEvent(new MeetingGroupProposalVerificationRequestedDomainEvent(this.Id));
        }

        private MeetingGroupProposal()
        {
            _decision = MeetingGroupProposalDecision.NoDecision;
        }

        public MeetingGroupProposalId Id { get; private set; }

        public void Accept(UserId userId)
        {
            this.CheckRule(new MeetingGroupProposalCanBeVerifiedOnceRule(_decision));

            _decision = MeetingGroupProposalDecision.AcceptDecision(DateTime.UtcNow, userId);

            _status = _decision.GetStatusForDecision();

            this.AddDomainEvent(new MeetingGroupProposalAcceptedDomainEvent(this.Id));
        }

        public void Reject(UserId userId, string rejectReason)
        {
            this.CheckRule(new MeetingGroupProposalCanBeVerifiedOnceRule(_decision));
            this.CheckRule(new MeetingGroupProposalRejectionMustHaveAReasonRule(rejectReason));

            _decision = MeetingGroupProposalDecision.RejectDecision(DateTime.UtcNow, userId, rejectReason);

            _status = _decision.GetStatusForDecision();

            this.AddDomainEvent(new MeetingGroupProposalRejectedDomainEvent(this.Id));
        }

        public static MeetingGroupProposal CreateToVerify(
            Guid meetingGroupProposalId,
            string name,
            string description,
            MeetingGroupLocation location,
            UserId proposalUserId,
            DateTime proposalDate)
        {
            var meetingGroupProposal = new MeetingGroupProposal(
                new MeetingGroupProposalId(meetingGroupProposalId),
                name,
                description,
                location,
                proposalUserId,
                proposalDate);

            return meetingGroupProposal;
        }
    }
}

*/

type MeetingGroupProposal struct {
	domain.Entity
	ID             MeetingGroupProposalsID
	Name           string
	Description    string
	Location       MeetingGroupLocation
	ProposalUserID users.UserID
	ProposalDate   time.Time
	Status         MeetingGroupProposalStatus
	Decision       MeetingGroupProposalDecision
}

// NewMeetingGroupProposal creates a new MeetingGroupProposal instance.
func NewMeetingGroupProposal(
	id MeetingGroupProposalsID,
	name, description string,
	location MeetingGroupLocation,
	proposalUserID users.UserID,
	proposalDate time.Time,
) *MeetingGroupProposal {
	return &MeetingGroupProposal{
		ID:             id,
		Name:           name,
		Description:    description,
		Location:       location,
		ProposalUserID: proposalUserID,
		ProposalDate:   proposalDate,
		Status:         MeetingGroupProposalStatusToVerify(),
		Decision:       MeetingGroupProposalDecisionNoDecision,
	}
}

// Accept marks the meeting group proposal as accepted.
func (m *MeetingGroupProposal) Accept(userID UserID) {
	m.CheckRule(MeetingGroupProposalCanBeVerifiedOnceRule(m.Decision))

	m.Decision = MeetingGroupProposalDecisionAcceptDecision(time.Now().UTC(), userID)
	m.Status = m.Decision.GetStatusForDecision()

	m.AddDomainEvent(MeetingGroupProposalAcceptedDomainEvent{ID: m.ID})
}

// Reject marks the meeting group proposal as rejected.
func (m *MeetingGroupProposal) Reject(userID UserID, rejectReason string) {
	m.CheckRule(MeetingGroupProposalCanBeVerifiedOnceRule(m.Decision))
	m.CheckRule(MeetingGroupProposalRejectionMustHaveAReasonRule(rejectReason))

	m.Decision = MeetingGroupProposalDecisionRejectDecision(time.Now().UTC(), userID, rejectReason)
	m.Status = m.Decision.GetStatusForDecision()

	m.AddDomainEvent(MeetingGroupProposalRejectedDomainEvent{ID: m.ID})
}

// CreateToMeetingGroupProposalsVerify creates a new meeting group proposal to be verified.
func CreateToMeetingGroupProposalsVerify(
	id MeetingGroupProposalsID,
	name, description string,
	location MeetingGroupLocation,
	proposalUserID UserID,
	proposalDate time.Time,
) *MeetingGroupProposal {
	meetingGroupProposal := NewMeetingGroupProposal(id, name, description, location, proposalUserID, proposalDate)

	meetingGroupProposal.AddDomainEvent(MeetingGroupProposalVerificationRequestedDomainEvent{ID: meetingGroupProposal.ID})

	return meetingGroupProposal
}
