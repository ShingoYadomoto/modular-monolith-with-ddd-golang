package domain

import (
	"time"

	"github.com/google/uuid"
)

/*
namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public class DomainEventBase : IDomainEvent
    {
        public Guid Id { get; }

        public DateTime OccurredOn { get; }

        public DomainEventBase()
        {
            this.Id = Guid.NewGuid();
            this.OccurredOn = DateTime.UtcNow;
        }
    }
}
*/

// DomainEventBase represents the base struct for domain events.
type DomainEventBase struct {
	id         uuid.UUID
	occurredOn time.Time
}

// NewDomainEventBase creates a new DomainEventBase instance.
func NewDomainEventBase() *DomainEventBase {
	return &DomainEventBase{
		id:         uuid.New(),
		occurredOn: time.Now(),
	}
}

func (d *DomainEventBase) ID() uuid.UUID {
	return d.id
}

func (d *DomainEventBase) OccurredOn() time.Time {
	return d.occurredOn
}
