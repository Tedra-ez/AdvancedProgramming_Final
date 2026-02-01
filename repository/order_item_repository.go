package repository

import (
	"context"

	"github.com/Tedra-ez/AdvancedProgramming_Final/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderItemStore interface {
	CreateMany(ctx context.Context, items []models.OrderItem) error
	FindByOrderId(ctx context.Context, orderID string) ([]*models.OrderItem, error)
}

type OrderItemRepositoryMongo struct {
	coll *mongo.Collection
}

func NewOrderItemRepositoryMongo(coll *mongo.Collection) *OrderItemRepositoryMongo {
	return &OrderItemRepositoryMongo{coll: coll}
}

func (r *OrderItemRepositoryMongo) CreateMany(ctx context.Context, items []models.OrderItem) error {
	if len(items) == 0 {
		return nil
	}
	docs := make([]interface{}, 0, len(items))
	for i := range items {
		doc := orderItemDocFromModel(&items[i])
		if items[i].ID == "" {
			id := primitive.NewObjectID()
			doc.ID = id
			items[i].ID = id.Hex()
		} else {
			oid, err := primitive.ObjectIDFromHex(items[i].ID)
			if err == nil {
				doc.ID = oid
			}
		}
		docs = append(docs, doc)
	}
	_, err := r.coll.InsertMany(ctx, docs)
	return err
}

func (r *OrderItemRepositoryMongo) FindByOrderId(ctx context.Context, orderID string) ([]*models.OrderItem, error) {
	cur, err := r.coll.Find(ctx, bson.M{"orderId": orderID}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []*models.OrderItem
	for cur.Next(ctx) {
		var doc orderItemDocStandalone
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		out = append(out, doc.toModel())
	}
	return out, cur.Err()
}

type orderItemDocStandalone struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OrderID       string             `bson:"orderId"`
	ProductID     string             `bson:"productId"`
	ProductName   string             `bson:"productName"`
	SelectedSize  string             `bson:"selectedSize"`
	SelectedColor string             `bson:"selectedColor"`
	Quantity      int                `bson:"quantity"`
	UnitPrice     float64            `bson:"unitPrice"`
	LineTotal     float64            `bson:"lineTotal"`
}

func orderItemDocFromModel(it *models.OrderItem) *orderItemDocStandalone {
	return &orderItemDocStandalone{
		OrderID:       it.OrderID,
		ProductID:     it.ProductID,
		ProductName:   it.ProductName,
		SelectedSize:  it.SelectedSize,
		SelectedColor: it.SelectedColor,
		Quantity:      it.Quantity,
		UnitPrice:     it.UnitPrice,
		LineTotal:     it.LineTotal,
	}
}

func (d *orderItemDocStandalone) toModel() *models.OrderItem {
	return &models.OrderItem{
		ID:            d.ID.Hex(),
		OrderID:       d.OrderID,
		ProductID:     d.ProductID,
		ProductName:   d.ProductName,
		SelectedSize:  d.SelectedSize,
		SelectedColor: d.SelectedColor,
		Quantity:      d.Quantity,
		UnitPrice:     d.UnitPrice,
		LineTotal:     d.LineTotal,
	}
}
