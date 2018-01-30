package base

import (
	"gonum.org/v1/gonum/mat"
)

// Classifier implementations predict categorical class labels.
type Classifier interface {
	// Takes a set of Instances, copies the class Attribute
	// and constructs a new set of Instances of equivalent
	// length with only the class Attribute and fills it in
	// with predictions.
	Predict(FixedDataGrid) (FixedDataGrid, error)
	// Takes a set of instances and updates the Classifier's
	// internal structures to enable prediction
	Fit(FixedDataGrid) error
	// Why not make every classifier return a nice-looking string?
	String() string
}

// BaseClassifier stores options common to every classifier.
type BaseClassifier struct {
	TrainingData *DataGrid
}

type BaseRegressor struct {
	Data   mat.Dense
	Name   string
	Labels []float64
}
