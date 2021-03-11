package bike

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewFirestoreRepository(client *firestore.Client) (firestoreRepository FirestoreRepository, err error) {
	firestoreRepository = FirestoreRepository{client: client}
	return
}

func (repo FirestoreRepository) GetByFilter(filter Filter) (bikes []Bike, err error) {
	var (
		bikesCollection = repo.client.Collection(colletionName)
		ctx             = context.Background()
	)

	bikes = []Bike{}

	if filter.Model != "" {
		bikesCollection.Query.Where("model", "==", filter.Model)
	}

	snapBikes, err := bikesCollection.Query.Documents(ctx).GetAll()
	if err != nil {
		return
	}

	for _, snapBike := range snapBikes {
		if filter.ID != "" && filter.ID != snapBike.Ref.ID {
			continue
		}

		bike := Bike{}

		bike, err = repo.documentSnapshotToBike(snapBike)
		if err != nil {
			return
		}

		bikes = append(bikes, bike)
	}

	return
}

func (repo FirestoreRepository) GetByID(bikeID string) (bike Bike, err error) {
	ctx := context.Background()

	snapBike, err := repo.client.Doc(bikeID).Get(ctx)
	if err != nil {
		return
	}

	bike, err = repo.documentSnapshotToBike(snapBike)
	return
}

func (repo FirestoreRepository) Register(bike Bike) (bikeID string, err error) {
	var (
		bikesCollection = repo.client.Collection(colletionName)
		ctx             = context.Background()
	)

	documentRef, _, err := bikesCollection.Add(ctx, bike)
	if err != nil {
		return
	}

	bikeID = documentRef.ID
	return
}

func (repo FirestoreRepository) Update(bike Bike) (err error) {
	var (
		bikesCollection = repo.client.Collection(colletionName)
		ctx             = context.Background()
	)

	bikeData := struct {
		Model string `firestore:"model"`
	}{
		Model: bike.Model,
	}

	_, err = bikesCollection.Doc(bike.ID).Set(ctx, bikeData)
	return
}

func (repo FirestoreRepository) Delete(bikeID string) (err error) {
	ctx := context.Background()
	_, err = repo.client.Collection(colletionName).Doc(bikeID).Delete(ctx)
	return
}

func (repo FirestoreRepository) documentSnapshotToBike(snapBike *firestore.DocumentSnapshot) (bike Bike, err error) {
	bikeData := struct {
		Model string `firestore:"model"`
	}{}

	err = snapBike.DataTo(&bikeData)
	if err != nil {
		return
	}

	bike.ID = snapBike.Ref.ID
	bike.Model = bikeData.Model
	return
}

func (repo FirestoreRepository) Start() (err error) {
	// todo: create bikes collection here
	return
}
