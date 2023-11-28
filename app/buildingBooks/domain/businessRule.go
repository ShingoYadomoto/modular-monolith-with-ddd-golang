package domain

/*
namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public interface IBusinessRule
    {
        bool IsBroken();

        string Message { get; }
    }
}
*/

type IBusinessRule interface {
	IsBroken() bool
	Message() string
}
