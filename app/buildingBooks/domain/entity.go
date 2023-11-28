package domain

/*
using System.Collections.Generic;

namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public abstract class Entity
    {
        private List<IDomainEvent> _domainEvents;

        /// <summary>
        /// Domain events occurred.
        /// </summary>
        public IReadOnlyCollection<IDomainEvent> DomainEvents => _domainEvents?.AsReadOnly();

        public void ClearDomainEvents()
        {
            _domainEvents?.Clear();
        }

        /// <summary>
        /// Add domain event.
        /// </summary>
        /// <param name="domainEvent">Domain event.</param>
        protected void AddDomainEvent(IDomainEvent domainEvent)
        {
            _domainEvents ??= new List<IDomainEvent>();

            this._domainEvents.Add(domainEvent);
        }

        protected void CheckRule(IBusinessRule rule)
        {
            if (rule.IsBroken())
            {
                throw new BusinessRuleValidationException(rule);
            }
        }
    }
}
*/

type Entity struct {
	domainEvents []IDomainEvent
}

// DomainEvents returns the domain events occurred.
func (e *Entity) DomainEvents() []IDomainEvent {
	return e.domainEvents
}

// ClearDomainEvents clears the domain events.
func (e *Entity) ClearDomainEvents() {
	e.domainEvents = nil
}

// AddDomainEvent adds a domain event.
func (e *Entity) AddDomainEvent(domainEvent IDomainEvent) {
	if e.domainEvents == nil {
		e.domainEvents = make([]IDomainEvent, 0)
	}

	e.domainEvents = append(e.domainEvents, domainEvent)
}

// CheckRule checks a business rule and throws an exception if it is broken.
func (e *Entity) CheckRule(rule IBusinessRule) {
	if rule.IsBroken() {
		panic(rule.Message())
	}
}
