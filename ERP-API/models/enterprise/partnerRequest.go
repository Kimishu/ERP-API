package enterprise

import (
	"ERP-API/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

/*
THIS STRUCTURE CREATED FOR GET/POST-REQUESTS
AND HANDLES MORE INFO FOR ENTERPRISES (GET),
			LESS INFO FOR ENTERPRISES (POST)
*/

type PartnerRequest struct {
	Id       uuid.UUID   `json:"id"`
	Sender   *Enterprise `json:"sender"`
	Reciever *Enterprise `json:"reciever"`
}

func (p *PartnerRequest) ReadIncoming(id string) []PartnerRequest {
	var incomingRequests []PartnerRequest

	rows, err := models.Database.Query(`SELECT req.id, 
													 s.id AS sender_id, 
													 s.name AS sender_name, 
													 s.email AS sender_email, 
													 s.subscription_id AS sender_subscription_id, 
													 sub.name AS sender_subscription_name,
													 e.id AS receiver_id, 
													 e.name AS receiver_name, 
													 e.email AS receiver_email, 
													 e.subscription_id AS receiver_subscription_id, 
													 sub2.name AS receiver_subscription_name
													 FROM 
														"PartnerRequests" req 
														JOIN "Enterprises" e ON req.reciever_id = e.id 
														JOIN "Enterprises" s ON req.sender_id = s.id 
														JOIN "Subscriptions" sub ON s.subscription_id = sub.id 
														JOIN "Subscriptions" sub2 ON e.subscription_id = sub2.id 
													 WHERE 
														req.reciever_id = $1`, id)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var request PartnerRequest
		var sender Enterprise
		var reciever Enterprise
		request.Sender = &sender
		request.Reciever = &reciever
		if err := rows.Scan(
			&request.Id,
			&sender.Id, &sender.Name, &sender.Email, &sender.Subscription.Id, &sender.Subscription.Name,
			&reciever.Id, &reciever.Name, &reciever.Email, &reciever.Subscription.Id, &reciever.Subscription.Name,
		); err != nil {
			fmt.Println(err)
			return []PartnerRequest{}
		}
		incomingRequests = append(incomingRequests, request)
	}
	return incomingRequests
}

func (p *PartnerRequest) ReadOutgoing(id string) []PartnerRequest {
	var outgoingRequests []PartnerRequest

	rows, err := models.Database.Query(`SELECT req.id, 
													 s.id AS sender_id, 
													 s.name AS sender_name, 
													 s.email AS sender_email, 
													 s.subscription_id AS sender_subscription_id, 
													 sub.name AS sender_subscription_name,
													 e.id AS receiver_id, 
													 e.name AS receiver_name, 
													 e.email AS receiver_email, 
													 e.subscription_id AS receiver_subscription_id, 
													 sub2.name AS receiver_subscription_name
													 FROM 
														"PartnerRequests" req 
														JOIN "Enterprises" e ON req.reciever_id = e.id 
														JOIN "Enterprises" s ON req.sender_id = s.id 
														JOIN "Subscriptions" sub ON s.subscription_id = sub.id 
														JOIN "Subscriptions" sub2 ON e.subscription_id = sub2.id 
													 WHERE 
														req.sender_id = $1`, id)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var request PartnerRequest
		var sender Enterprise
		var reciever Enterprise
		request.Reciever = &reciever
		request.Sender = &sender
		if err := rows.Scan(
			&request.Id,
			&sender.Id, &sender.Name, &sender.Email, &sender.Subscription.Id, &sender.Subscription.Name,
			&reciever.Id, &reciever.Name, &reciever.Email, &reciever.Subscription.Id, &reciever.Subscription.Name,
		); err != nil {
			fmt.Println(err)
			return []PartnerRequest{}
		}
		outgoingRequests = append(outgoingRequests, request)
	}
	return outgoingRequests
}

func (p *PartnerRequest) Write() error {
	_, err := models.Database.Exec("INSERT INTO \"PartnerRequests\" VALUES ($1, $2, $3)", uuid.New(), p.Sender.Id, p.Reciever.Id)
	if err != nil {
		return errors.New("Failed to create a new partner request!")
	}
	return nil
}

func (p *PartnerRequest) Accept(recieverId string) error {
	if p.Sender.Id.String() == recieverId {
		return errors.New("Sender can't accept request :)")
	}

	_, err := models.Database.Exec("INSERT INTO \"Partners\" VALUES ($1, $2)", p.Sender.Id, p.Reciever.Id)
	if err != nil {
		return err
	}
	err = p.Decline(recieverId)
	if err != nil {
		return err
	}

	return nil
}

func (p *PartnerRequest) Decline(enterpriseId string) error {
	_, err := models.Database.Exec("DELETE FROM \"PartnerRequests\" WHERE id = $1", p.Id)
	if err != nil {
		return errors.New("Rejection error")
	}

	return nil
}
