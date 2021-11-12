package hedera

import (
	"time"

	"github.com/hashgraph/hedera-protobufs-go/services"
)

type TopicInfoQuery struct {
	Query
	topicID *TopicID
}

// NewTopicInfoQuery creates a TopicInfoQuery query which can be used to construct and execute a
//  Get Topic Info Query.
func NewTopicInfoQuery() *TopicInfoQuery {
	return &TopicInfoQuery{
		Query: _NewQuery(true),
	}
}

// SetTopicID sets the topic to retrieve info about (the parameters and running state of).
func (query *TopicInfoQuery) SetTopicID(topicID TopicID) *TopicInfoQuery {
	query.topicID = &topicID
	return query
}

func (query *TopicInfoQuery) GetTopicID() TopicID {
	if query.topicID == nil {
		return TopicID{}
	}

	return *query.topicID
}

func (query *TopicInfoQuery) _ValidateNetworkOnIDs(client *Client) error {
	if client == nil || !client.autoValidateChecksums {
		return nil
	}

	if query.topicID != nil {
		if err := query.topicID.ValidateChecksum(client); err != nil {
			return err
		}
	}

	return nil
}

func (query *TopicInfoQuery) _Build() *services.Query_ConsensusGetTopicInfo {
	body := &services.ConsensusGetTopicInfoQuery{
		Header: &services.QueryHeader{},
	}

	if query.topicID != nil {
		body.TopicID = query.topicID._ToProtobuf()
	}

	return &services.Query_ConsensusGetTopicInfo{
		ConsensusGetTopicInfo: body,
	}
}

func (query *TopicInfoQuery) _QueryMakeRequest() _ProtoRequest {
	pb := query._Build()
	if query.isPaymentRequired && len(query.paymentTransactions) > 0 {
		pb.ConsensusGetTopicInfo.Header.Payment = query.paymentTransactions[query.nextPaymentTransactionIndex]
	}
	pb.ConsensusGetTopicInfo.Header.ResponseType = services.ResponseType_ANSWER_ONLY

	return _ProtoRequest{
		query: &services.Query{
			Query: pb,
		},
	}
}

func (query *TopicInfoQuery) _CostQueryMakeRequest(client *Client) (_ProtoRequest, error) {
	pb := query._Build()

	paymentTransaction, err := _QueryMakePaymentTransaction(TransactionID{}, AccountID{}, client.operator, Hbar{})
	if err != nil {
		return _ProtoRequest{}, err
	}

	pb.ConsensusGetTopicInfo.Header.Payment = paymentTransaction
	pb.ConsensusGetTopicInfo.Header.ResponseType = services.ResponseType_COST_ANSWER

	return _ProtoRequest{
		query: &services.Query{
			Query: pb,
		},
	}, nil
}

func (query *TopicInfoQuery) GetCost(client *Client) (Hbar, error) {
	if client == nil || client.operator == nil {
		return Hbar{}, errNoClientProvided
	}

	query.nodeIDs = client.network._GetNodeAccountIDsForExecute()

	err := query._ValidateNetworkOnIDs(client)
	if err != nil {
		return Hbar{}, err
	}

	protoReq, err := query._CostQueryMakeRequest(client)
	if err != nil {
		return Hbar{}, err
	}

	resp, err := _Execute(
		client,
		_Request{
			query: &query.Query,
		},
		_TopicInfoQueryShouldRetry,
		protoReq,
		_CostQueryAdvanceRequest,
		_CostQueryGetNodeAccountID,
		_TopicInfoQueryGetMethod,
		_TopicInfoQueryMapStatusError,
		_QueryMapResponse,
	)

	if err != nil {
		return Hbar{}, err
	}

	cost := int64(resp.query.GetConsensusGetTopicInfo().Header.Cost)
	if cost < 25 {
		return HbarFromTinybar(25), nil
	}
	return HbarFromTinybar(cost), nil
}

func _TopicInfoQueryShouldRetry(_ _Request, response _Response) _ExecutionState {
	return _QueryShouldRetry(Status(response.query.GetConsensusGetTopicInfo().Header.NodeTransactionPrecheckCode))
}

func _TopicInfoQueryMapStatusError(_ _Request, response _Response) error {
	return ErrHederaPreCheckStatus{
		Status: Status(response.query.GetConsensusGetTopicInfo().Header.NodeTransactionPrecheckCode),
	}
}

func _TopicInfoQueryGetMethod(_ _Request, channel *_Channel) _Method {
	return _Method{
		query: channel._GetTopic().GetTopicInfo,
	}
}

// Execute executes the TopicInfoQuery using the provided client
func (query *TopicInfoQuery) Execute(client *Client) (TopicInfo, error) {
	if client == nil || client.operator == nil {
		return TopicInfo{}, errNoClientProvided
	}

	if len(query.Query.GetNodeAccountIDs()) == 0 {
		query.SetNodeAccountIDs(client.network._GetNodeAccountIDsForExecute())
	}

	err := query._ValidateNetworkOnIDs(client)
	if err != nil {
		return TopicInfo{}, err
	}

	query.paymentTransactionID = TransactionIDGenerate(client.operator.accountID)

	var cost Hbar
	if query.queryPayment.tinybar != 0 {
		cost = query.queryPayment
	} else {
		if query.maxQueryPayment.tinybar == 0 {
			cost = client.maxQueryPayment
		} else {
			cost = query.maxQueryPayment
		}

		actualCost, err := query.GetCost(client)
		if err != nil {
			return TopicInfo{}, err
		}

		if cost.tinybar < actualCost.tinybar {
			return TopicInfo{}, ErrMaxQueryPaymentExceeded{
				QueryCost:       actualCost,
				MaxQueryPayment: cost,
				query:           "TopicInfoQuery",
			}
		}

		cost = actualCost
	}

	err = _QueryGeneratePayments(&query.Query, client, cost)
	if err != nil {
		return TopicInfo{}, err
	}

	resp, err := _Execute(
		client,
		_Request{
			query: &query.Query,
		},
		_TopicInfoQueryShouldRetry,
		query._QueryMakeRequest(),
		_QueryAdvanceRequest,
		_QueryGetNodeAccountID,
		_TopicInfoQueryGetMethod,
		_TopicInfoQueryMapStatusError,
		_QueryMapResponse,
	)

	if err != nil {
		return TopicInfo{}, err
	}

	return _TopicInfoFromProtobuf(resp.query.GetConsensusGetTopicInfo().TopicInfo)
}

// SetMaxQueryPayment sets the maximum payment allowed for this Query.
func (query *TopicInfoQuery) SetMaxQueryPayment(maxPayment Hbar) *TopicInfoQuery {
	query.Query.SetMaxQueryPayment(maxPayment)
	return query
}

// SetQueryPayment sets the payment amount for this Query.
func (query *TopicInfoQuery) SetQueryPayment(paymentAmount Hbar) *TopicInfoQuery {
	query.Query.SetQueryPayment(paymentAmount)
	return query
}

func (query *TopicInfoQuery) SetNodeAccountIDs(accountID []AccountID) *TopicInfoQuery {
	query.Query.SetNodeAccountIDs(accountID)
	return query
}

func (query *TopicInfoQuery) SetMaxRetry(count int) *TopicInfoQuery {
	query.Query.SetMaxRetry(count)
	return query
}

func (query *TopicInfoQuery) SetMaxBackoff(max time.Duration) *TopicInfoQuery {
	if max.Nanoseconds() < 0 {
		panic("maxBackoff must be a positive duration")
	} else if max.Nanoseconds() < query.minBackoff.Nanoseconds() {
		panic("maxBackoff must be greater than or equal to minBackoff")
	}
	query.maxBackoff = &max
	return query
}

func (query *TopicInfoQuery) GetMaxBackoff() time.Duration {
	if query.maxBackoff != nil {
		return *query.maxBackoff
	}

	return 8 * time.Second
}

func (query *TopicInfoQuery) SetMinBackoff(min time.Duration) *TopicInfoQuery {
	if min.Nanoseconds() < 0 {
		panic("minBackoff must be a positive duration")
	} else if query.maxBackoff.Nanoseconds() < min.Nanoseconds() {
		panic("minBackoff must be less than or equal to maxBackoff")
	}
	query.minBackoff = &min
	return query
}

func (query *TopicInfoQuery) GetMinBackoff() time.Duration {
	if query.minBackoff != nil {
		return *query.minBackoff
	}

	return 250 * time.Millisecond
}
