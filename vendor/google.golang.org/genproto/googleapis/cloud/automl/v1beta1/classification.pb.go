// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/classification.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type of the classification problem.
type ClassificationType int32

const (
	// An un-set value of this enum.
	ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED ClassificationType = 0
	// At most one label is allowed per example.
	ClassificationType_MULTICLASS ClassificationType = 1
	// Multiple labels are allowed for one example.
	ClassificationType_MULTILABEL ClassificationType = 2
)

var ClassificationType_name = map[int32]string{
	0: "CLASSIFICATION_TYPE_UNSPECIFIED",
	1: "MULTICLASS",
	2: "MULTILABEL",
}

var ClassificationType_value = map[string]int32{
	"CLASSIFICATION_TYPE_UNSPECIFIED": 0,
	"MULTICLASS":                      1,
	"MULTILABEL":                      2,
}

func (x ClassificationType) String() string {
	return proto.EnumName(ClassificationType_name, int32(x))
}

func (ClassificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{0}
}

// Contains annotation details specific to classification.
type ClassificationAnnotation struct {
	// Output only. A confidence estimate between 0.0 and 1.0. A higher value
	// means greater confidence that the annotation is positive. If a user
	// approves an annotation as negative or positive, the score value remains
	// unchanged. If a user creates an annotation, the score is 0 for negative or
	// 1 for positive.
	Score                float32  `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationAnnotation) Reset()         { *m = ClassificationAnnotation{} }
func (m *ClassificationAnnotation) String() string { return proto.CompactTextString(m) }
func (*ClassificationAnnotation) ProtoMessage()    {}
func (*ClassificationAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{0}
}

func (m *ClassificationAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationAnnotation.Unmarshal(m, b)
}
func (m *ClassificationAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationAnnotation.Marshal(b, m, deterministic)
}
func (m *ClassificationAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationAnnotation.Merge(m, src)
}
func (m *ClassificationAnnotation) XXX_Size() int {
	return xxx_messageInfo_ClassificationAnnotation.Size(m)
}
func (m *ClassificationAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationAnnotation proto.InternalMessageInfo

func (m *ClassificationAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Contains annotation details specific to video classification.
type VideoClassificationAnnotation struct {
	// Output only. Expresses the type of video classification. Possible values:
	//
	// *  `segment` - Classification done on a specified by user
	//        time segment of a video. AnnotationSpec is answered to be present
	//        in that time segment, if it is present in any part of it. The video
	//        ML model evaluations are done only for this type of classification.
	//
	// *  `shot`- Shot-level classification.
	//        AutoML Video Intelligence determines the boundaries
	//        for each camera shot in the entire segment of the video that user
	//        specified in the request configuration. AutoML Video Intelligence
	//        then returns labels and their confidence scores for each detected
	//        shot, along with the start and end time of the shot.
	//        WARNING: Model evaluation is not done for this classification type,
	//        the quality of it depends on training data, but there are no
	//        metrics provided to describe that quality.
	//
	// *  `1s_interval` - AutoML Video Intelligence returns labels and their
	//        confidence scores for each second of the entire segment of the video
	//        that user specified in the request configuration.
	//        WARNING: Model evaluation is not done for this classification type,
	//        the quality of it depends on training data, but there are no
	//        metrics provided to describe that quality.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Output only . The classification details of this annotation.
	ClassificationAnnotation *ClassificationAnnotation `protobuf:"bytes,2,opt,name=classification_annotation,json=classificationAnnotation,proto3" json:"classification_annotation,omitempty"`
	// Output only . The time segment of the video to which the
	// annotation applies.
	TimeSegment          *TimeSegment `protobuf:"bytes,3,opt,name=time_segment,json=timeSegment,proto3" json:"time_segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VideoClassificationAnnotation) Reset()         { *m = VideoClassificationAnnotation{} }
func (m *VideoClassificationAnnotation) String() string { return proto.CompactTextString(m) }
func (*VideoClassificationAnnotation) ProtoMessage()    {}
func (*VideoClassificationAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{1}
}

func (m *VideoClassificationAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationAnnotation.Unmarshal(m, b)
}
func (m *VideoClassificationAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationAnnotation.Marshal(b, m, deterministic)
}
func (m *VideoClassificationAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationAnnotation.Merge(m, src)
}
func (m *VideoClassificationAnnotation) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationAnnotation.Size(m)
}
func (m *VideoClassificationAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationAnnotation proto.InternalMessageInfo

func (m *VideoClassificationAnnotation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VideoClassificationAnnotation) GetClassificationAnnotation() *ClassificationAnnotation {
	if m != nil {
		return m.ClassificationAnnotation
	}
	return nil
}

func (m *VideoClassificationAnnotation) GetTimeSegment() *TimeSegment {
	if m != nil {
		return m.TimeSegment
	}
	return nil
}

// Model evaluation metrics for classification problems.
// Note: For Video Classification this metrics only describe quality of the
// Video Classification predictions of "segment_classification" type.
type ClassificationEvaluationMetrics struct {
	// Output only. The Area Under Precision-Recall Curve metric. Micro-averaged
	// for the overall evaluation.
	AuPrc float32 `protobuf:"fixed32,1,opt,name=au_prc,json=auPrc,proto3" json:"au_prc,omitempty"`
	// Output only. The Area Under Precision-Recall Curve metric based on priors.
	// Micro-averaged for the overall evaluation.
	// Deprecated.
	BaseAuPrc float32 `protobuf:"fixed32,2,opt,name=base_au_prc,json=baseAuPrc,proto3" json:"base_au_prc,omitempty"` // Deprecated: Do not use.
	// Output only. The Area Under Receiver Operating Characteristic curve metric.
	// Micro-averaged for the overall evaluation.
	AuRoc float32 `protobuf:"fixed32,6,opt,name=au_roc,json=auRoc,proto3" json:"au_roc,omitempty"`
	// Output only. The Log Loss metric.
	LogLoss float32 `protobuf:"fixed32,7,opt,name=log_loss,json=logLoss,proto3" json:"log_loss,omitempty"`
	// Output only. Metrics for each confidence_threshold in
	// 0.00,0.05,0.10,...,0.95,0.96,0.97,0.98,0.99 and
	// position_threshold = INT32_MAX_VALUE.
	// ROC and precision-recall curves, and other aggregated metrics are derived
	// from them. The confidence metrics entries may also be supplied for
	// additional values of position_threshold, but from these no aggregated
	// metrics are computed.
	ConfidenceMetricsEntry []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entry,json=confidenceMetricsEntry,proto3" json:"confidence_metrics_entry,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for MULTICLASS classification problems where number
	// of labels is no more than 10.
	// Only set for model level evaluation, not for evaluation per label.
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,4,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	// Output only. The annotation spec ids used for this evaluation.
	AnnotationSpecId     []string `protobuf:"bytes,5,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics) Reset()         { *m = ClassificationEvaluationMetrics{} }
func (m *ClassificationEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*ClassificationEvaluationMetrics) ProtoMessage()    {}
func (*ClassificationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{2}
}

func (m *ClassificationEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics.Size(m)
}
func (m *ClassificationEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics) GetAuPrc() float32 {
	if m != nil {
		return m.AuPrc
	}
	return 0
}

// Deprecated: Do not use.
func (m *ClassificationEvaluationMetrics) GetBaseAuPrc() float32 {
	if m != nil {
		return m.BaseAuPrc
	}
	return 0
}

func (m *ClassificationEvaluationMetrics) GetAuRoc() float32 {
	if m != nil {
		return m.AuRoc
	}
	return 0
}

func (m *ClassificationEvaluationMetrics) GetLogLoss() float32 {
	if m != nil {
		return m.LogLoss
	}
	return 0
}

func (m *ClassificationEvaluationMetrics) GetConfidenceMetricsEntry() []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry {
	if m != nil {
		return m.ConfidenceMetricsEntry
	}
	return nil
}

func (m *ClassificationEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if m != nil {
		return m.ConfusionMatrix
	}
	return nil
}

func (m *ClassificationEvaluationMetrics) GetAnnotationSpecId() []string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return nil
}

// Metrics for a single confidence threshold.
type ClassificationEvaluationMetrics_ConfidenceMetricsEntry struct {
	// Output only. Metrics are computed with an assumption that the model
	// never returns predictions with score lower than this value.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Metrics are computed with an assumption that the model
	// always returns at most this many predictions (ordered by their score,
	// descendingly), but they all still need to meet the confidence_threshold.
	PositionThreshold int32 `protobuf:"varint,14,opt,name=position_threshold,json=positionThreshold,proto3" json:"position_threshold,omitempty"`
	// Output only. Recall (True Positive Rate) for the given confidence
	// threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision for the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. False Positive Rate for the given confidence threshold.
	FalsePositiveRate float32 `protobuf:"fixed32,8,opt,name=false_positive_rate,json=falsePositiveRate,proto3" json:"false_positive_rate,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. The Recall (True Positive Rate) when only considering the
	// label that has the highest prediction score and not below the confidence
	// threshold for each example.
	RecallAt1 float32 `protobuf:"fixed32,5,opt,name=recall_at1,json=recallAt1,proto3" json:"recall_at1,omitempty"`
	// Output only. The precision when only considering the label that has the
	// highest prediction score and not below the confidence threshold for each
	// example.
	PrecisionAt1 float32 `protobuf:"fixed32,6,opt,name=precision_at1,json=precisionAt1,proto3" json:"precision_at1,omitempty"`
	// Output only. The False Positive Rate when only considering the label that
	// has the highest prediction score and not below the confidence threshold
	// for each example.
	FalsePositiveRateAt1 float32 `protobuf:"fixed32,9,opt,name=false_positive_rate_at1,json=falsePositiveRateAt1,proto3" json:"false_positive_rate_at1,omitempty"`
	// Output only. The harmonic mean of [recall_at1][google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision_at1].
	F1ScoreAt1 float32 `protobuf:"fixed32,7,opt,name=f1_score_at1,json=f1ScoreAt1,proto3" json:"f1_score_at1,omitempty"`
	// Output only. The number of model created labels that match a ground truth
	// label.
	TruePositiveCount int64 `protobuf:"varint,10,opt,name=true_positive_count,json=truePositiveCount,proto3" json:"true_positive_count,omitempty"`
	// Output only. The number of model created labels that do not match a
	// ground truth label.
	FalsePositiveCount int64 `protobuf:"varint,11,opt,name=false_positive_count,json=falsePositiveCount,proto3" json:"false_positive_count,omitempty"`
	// Output only. The number of ground truth labels that are not matched
	// by a model created label.
	FalseNegativeCount int64 `protobuf:"varint,12,opt,name=false_negative_count,json=falseNegativeCount,proto3" json:"false_negative_count,omitempty"`
	// Output only. The number of labels that were not created by the model,
	// but if they would, they would not match a ground truth label.
	TrueNegativeCount    int64    `protobuf:"varint,13,opt,name=true_negative_count,json=trueNegativeCount,proto3" json:"true_negative_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Reset() {
	*m = ClassificationEvaluationMetrics_ConfidenceMetricsEntry{}
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{2, 0}
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfidenceMetricsEntry proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if m != nil {
		return m.ConfidenceThreshold
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPositionThreshold() int32 {
	if m != nil {
		return m.PositionThreshold
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveRate() float32 {
	if m != nil {
		return m.FalsePositiveRate
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecallAt1() float32 {
	if m != nil {
		return m.RecallAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecisionAt1() float32 {
	if m != nil {
		return m.PrecisionAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveRateAt1() float32 {
	if m != nil {
		return m.FalsePositiveRateAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1ScoreAt1() float32 {
	if m != nil {
		return m.F1ScoreAt1
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetTruePositiveCount() int64 {
	if m != nil {
		return m.TruePositiveCount
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveCount() int64 {
	if m != nil {
		return m.FalsePositiveCount
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalseNegativeCount() int64 {
	if m != nil {
		return m.FalseNegativeCount
	}
	return 0
}

func (m *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetTrueNegativeCount() int64 {
	if m != nil {
		return m.TrueNegativeCount
	}
	return 0
}

// Confusion matrix of the model running the classification.
type ClassificationEvaluationMetrics_ConfusionMatrix struct {
	// Output only. IDs of the annotation specs used in the confusion matrix.
	// For Tables CLASSIFICATION
	//
	// [prediction_type][google.cloud.automl.v1beta1.TablesModelMetadata.prediction_type]
	// only list of [annotation_spec_display_name-s][] is populated.
	AnnotationSpecId []string `protobuf:"bytes,1,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. Rows in the confusion matrix. The number of rows is equal to
	// the size of `annotation_spec_id`.
	// `row[i].value[j]` is the number of examples that have ground truth of the
	// `annotation_spec_id[i]` and are predicted as `annotation_spec_id[j]` by
	// the model being evaluated.
	Row                  []*ClassificationEvaluationMetrics_ConfusionMatrix_Row `protobuf:"bytes,2,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) Reset() {
	*m = ClassificationEvaluationMetrics_ConfusionMatrix{}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfusionMatrix) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfusionMatrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{2, 1}
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) GetAnnotationSpecId() []string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return nil
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix) GetRow() []*ClassificationEvaluationMetrics_ConfusionMatrix_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

// Output only. A row in the confusion matrix.
type ClassificationEvaluationMetrics_ConfusionMatrix_Row struct {
	// Output only. Value of the specific cell in the confusion matrix.
	// The number of values each row has (i.e. the length of the row) is equal
	// to the length of the `annotation_spec_id` field or, if that one is not
	// populated, length of the [display_name][google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix.display_name] field.
	ExampleCount         []int32  `protobuf:"varint,1,rep,packed,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) Reset() {
	*m = ClassificationEvaluationMetrics_ConfusionMatrix_Row{}
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) String() string {
	return proto.CompactTextString(m)
}
func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) ProtoMessage() {}
func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b436fefe6ae5367, []int{2, 1, 0}
}

func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Unmarshal(m, b)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Marshal(b, m, deterministic)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Merge(m, src)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_Size() int {
	return xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.Size(m)
}
func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationEvaluationMetrics_ConfusionMatrix_Row proto.InternalMessageInfo

func (m *ClassificationEvaluationMetrics_ConfusionMatrix_Row) GetExampleCount() []int32 {
	if m != nil {
		return m.ExampleCount
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.automl.v1beta1.ClassificationType", ClassificationType_name, ClassificationType_value)
	proto.RegisterType((*ClassificationAnnotation)(nil), "google.cloud.automl.v1beta1.ClassificationAnnotation")
	proto.RegisterType((*VideoClassificationAnnotation)(nil), "google.cloud.automl.v1beta1.VideoClassificationAnnotation")
	proto.RegisterType((*ClassificationEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfidenceMetricsEntry)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfusionMatrix)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix")
	proto.RegisterType((*ClassificationEvaluationMetrics_ConfusionMatrix_Row)(nil), "google.cloud.automl.v1beta1.ClassificationEvaluationMetrics.ConfusionMatrix.Row")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/classification.proto", fileDescriptor_7b436fefe6ae5367)
}

var fileDescriptor_7b436fefe6ae5367 = []byte{
	// 859 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0x9d, 0xec, 0x38, 0xad, 0xaf, 0xdd, 0xd6, 0x61, 0xb2, 0x4c, 0x75, 0xdb, 0xc5, 0x48, 0x5f,
	0x8c, 0x60, 0x93, 0xe3, 0x0e, 0x7d, 0xf1, 0x9e, 0x1c, 0xcf, 0x1d, 0x8c, 0x39, 0x99, 0x41, 0xbb,
	0x05, 0x3a, 0x04, 0x10, 0x18, 0x9a, 0x56, 0x05, 0x48, 0xa2, 0x40, 0x52, 0x49, 0xfb, 0x11, 0xfb,
	0x9f, 0x7d, 0xc3, 0x5e, 0xf6, 0x03, 0xfb, 0x82, 0x3d, 0x0f, 0x7b, 0x1e, 0x48, 0xca, 0x56, 0x9c,
	0x38, 0xc1, 0x86, 0xed, 0xcd, 0xbc, 0xe7, 0x9c, 0x7b, 0xce, 0xa5, 0xae, 0x2c, 0x38, 0x0e, 0x38,
	0x0f, 0x22, 0xd6, 0xa1, 0x11, 0xcf, 0xe6, 0x1d, 0x92, 0x29, 0x1e, 0x47, 0x9d, 0xcb, 0xee, 0x05,
	0x53, 0xa4, 0xdb, 0xa1, 0x11, 0x91, 0x32, 0x5c, 0x84, 0x94, 0xa8, 0x90, 0x27, 0x5e, 0x2a, 0xb8,
	0xe2, 0xe8, 0x99, 0x55, 0x78, 0x46, 0xe1, 0x59, 0x85, 0x97, 0x2b, 0x9a, 0x47, 0xf7, 0xb5, 0x53,
	0x2c, 0x4e, 0xb9, 0x20, 0x91, 0x6d, 0xd4, 0x7c, 0x9e, 0x73, 0x49, 0x1a, 0x76, 0x48, 0x92, 0x70,
	0x65, 0x5c, 0xa4, 0x45, 0x0f, 0x8f, 0xc1, 0x1d, 0xac, 0xd9, 0xf7, 0x57, 0x14, 0xb4, 0x07, 0x15,
	0x49, 0xb9, 0x60, 0xae, 0xd3, 0x72, 0xda, 0x25, 0x6c, 0x0f, 0x87, 0x7f, 0x3a, 0xf0, 0xe2, 0x5d,
	0x38, 0x67, 0xfc, 0x4e, 0x1d, 0x82, 0x2d, 0xf5, 0x29, 0xb5, 0xb2, 0x2a, 0x36, 0xbf, 0x91, 0x80,
	0xa7, 0xeb, 0x63, 0xfa, 0x45, 0x16, 0xb7, 0xd4, 0x72, 0xda, 0xb5, 0x57, 0xaf, 0xbd, 0x7b, 0x46,
	0xf6, 0xee, 0x72, 0xc3, 0x2e, 0xbd, 0x2b, 0xc7, 0x0f, 0x50, 0x57, 0x61, 0xcc, 0x7c, 0xc9, 0x82,
	0x98, 0x25, 0xca, 0x2d, 0x1b, 0x9b, 0xf6, 0xbd, 0x36, 0xb3, 0x30, 0x66, 0x53, 0xcb, 0xc7, 0x35,
	0x55, 0x1c, 0x0e, 0xff, 0xaa, 0xc2, 0xc1, 0x7a, 0x86, 0xe1, 0x25, 0x89, 0x32, 0xf3, 0xeb, 0x94,
	0x29, 0x11, 0x52, 0x89, 0x3e, 0x87, 0x6d, 0x92, 0xf9, 0xa9, 0xa0, 0xcb, 0x1b, 0x23, 0xd9, 0x44,
	0x50, 0x74, 0x08, 0xb5, 0x0b, 0x22, 0x99, 0x9f, 0x63, 0x7a, 0xda, 0xd2, 0x49, 0xc9, 0x75, 0x70,
	0x55, 0x97, 0xfb, 0x86, 0x63, 0xa5, 0x82, 0x53, 0x77, 0x7b, 0x29, 0xc5, 0x9c, 0xa2, 0xa7, 0xf0,
	0x30, 0xe2, 0x81, 0x1f, 0x71, 0x29, 0xdd, 0x07, 0x06, 0x78, 0x10, 0xf1, 0x60, 0xcc, 0xa5, 0x44,
	0x3f, 0x3b, 0xe0, 0x52, 0x9e, 0x2c, 0xc2, 0x39, 0x4b, 0x28, 0xf3, 0x63, 0x9b, 0xc1, 0x67, 0x89,
	0x12, 0x9f, 0xdc, 0x72, 0xab, 0xdc, 0xae, 0xbd, 0x9a, 0xfe, 0x8b, 0x1b, 0xbd, 0x35, 0x8d, 0x37,
	0x58, 0x35, 0xcf, 0x2b, 0x43, 0xdd, 0x1a, 0xef, 0xd3, 0x8d, 0x75, 0x74, 0x05, 0x0d, 0x8d, 0x64,
	0x52, 0x3f, 0xdc, 0x98, 0x28, 0x11, 0x7e, 0x74, 0xb7, 0xcc, 0x8d, 0x8f, 0xff, 0x73, 0x0c, 0xd3,
	0xf4, 0xd4, 0xf4, 0xc4, 0x4f, 0xe8, 0x7a, 0x01, 0x7d, 0x05, 0xa8, 0xd8, 0x25, 0x5f, 0xa6, 0x8c,
	0xfa, 0xe1, 0xdc, 0xad, 0xb4, 0xca, 0xed, 0x2a, 0x6e, 0x14, 0xc8, 0x34, 0x65, 0x74, 0x34, 0x6f,
	0xfe, 0xbe, 0x05, 0xfb, 0x9b, 0x27, 0x43, 0x5d, 0xd8, 0xbb, 0x76, 0xa1, 0xea, 0x83, 0x60, 0xf2,
	0x03, 0x8f, 0xe6, 0xf9, 0xc3, 0xdc, 0x2d, 0xb0, 0xd9, 0x12, 0x42, 0x5f, 0x03, 0x4a, 0xb9, 0x0c,
	0x8d, 0x73, 0x21, 0x78, 0xdc, 0x72, 0xda, 0x15, 0xbc, 0xb3, 0x44, 0x0a, 0xfa, 0x3e, 0x6c, 0x0b,
	0x46, 0x49, 0x14, 0xd9, 0x25, 0xc0, 0xf9, 0x09, 0x3d, 0x87, 0x6a, 0x2a, 0x18, 0x0d, 0xf5, 0x54,
	0x66, 0x4d, 0x4b, 0xb8, 0x28, 0x20, 0x0f, 0x76, 0x17, 0x24, 0x92, 0xcc, 0xb7, 0x0d, 0x2f, 0x99,
	0x2f, 0x88, 0x62, 0xee, 0x43, 0xc3, 0xdb, 0x31, 0xd0, 0x24, 0x47, 0x30, 0x51, 0x4c, 0x2f, 0xcd,
	0xa2, 0xeb, 0xdb, 0x57, 0x77, 0xcb, 0x2e, 0xcd, 0xa2, 0x3b, 0xd5, 0x47, 0xf4, 0x02, 0xc0, 0x5a,
	0xfa, 0x44, 0x75, 0xdd, 0x8a, 0x75, 0xb2, 0x95, 0xbe, 0xea, 0xa2, 0x97, 0xf0, 0x68, 0x65, 0x6b,
	0x18, 0x76, 0x19, 0xeb, 0xab, 0xa2, 0x26, 0xbd, 0x86, 0x2f, 0x36, 0xc4, 0x31, 0xf4, 0xaa, 0xa1,
	0xef, 0xdd, 0x8a, 0xa4, 0x65, 0x2d, 0xa8, 0x2f, 0x53, 0x19, 0xae, 0x5d, 0x67, 0xc8, 0x93, 0x69,
	0x86, 0x07, 0xbb, 0x4a, 0x64, 0xd7, 0xfa, 0x52, 0x9e, 0x25, 0xca, 0x85, 0x96, 0xd3, 0x2e, 0xe3,
	0x1d, 0x0d, 0x2d, 0x7b, 0x0e, 0x34, 0x80, 0x8e, 0x61, 0xef, 0x46, 0x10, 0x2b, 0xa8, 0x19, 0x01,
	0x5a, 0x4b, 0x71, 0x43, 0x91, 0xb0, 0x80, 0x5c, 0x53, 0xd4, 0xaf, 0x29, 0xce, 0x72, 0xc8, 0x2a,
	0x96, 0x99, 0x6e, 0x08, 0x1e, 0x15, 0x99, 0xd6, 0xf8, 0xcd, 0xdf, 0x1c, 0x78, 0x32, 0xf8, 0x47,
	0x0b, 0xea, 0x6c, 0x5e, 0x50, 0x74, 0x01, 0x65, 0xc1, 0xaf, 0xdc, 0x92, 0x79, 0x83, 0x27, 0xff,
	0xe7, 0xab, 0xe3, 0x61, 0x7e, 0x85, 0x75, 0xf3, 0xe6, 0x11, 0x94, 0x31, 0xbf, 0xd2, 0x8f, 0x9b,
	0x7d, 0x24, 0x71, 0x1a, 0x2d, 0xc7, 0xd2, 0x99, 0x2a, 0xb8, 0x9e, 0x17, 0xcd, 0x44, 0x47, 0xef,
	0x01, 0xad, 0xfb, 0xcc, 0xf4, 0xff, 0xf9, 0x4b, 0x38, 0x18, 0x8c, 0xfb, 0xd3, 0xe9, 0xe8, 0xcd,
	0x68, 0xd0, 0x9f, 0x8d, 0x7e, 0x3c, 0xf3, 0x67, 0xef, 0x27, 0x43, 0xff, 0xed, 0xd9, 0x74, 0x32,
	0x1c, 0x8c, 0xde, 0x8c, 0x86, 0xdf, 0x35, 0x3e, 0x43, 0x8f, 0x01, 0x4e, 0xdf, 0x8e, 0x67, 0x23,
	0xc3, 0x6c, 0x38, 0xab, 0xf3, 0xb8, 0x7f, 0x32, 0x1c, 0x37, 0x4a, 0x27, 0xbf, 0x38, 0x70, 0x40,
	0x79, 0x7c, 0xdf, 0x8c, 0x27, 0xbb, 0xeb, 0xe6, 0x13, 0xfd, 0xd5, 0xfa, 0xa9, 0x9f, 0x2b, 0x02,
	0x1e, 0x91, 0x24, 0xf0, 0xb8, 0x08, 0x3a, 0x01, 0x4b, 0xcc, 0x17, 0xad, 0x63, 0x21, 0x92, 0x86,
	0x72, 0xe3, 0xe7, 0xf1, 0x5b, 0x7b, 0xfc, 0xb5, 0xf4, 0xec, 0x7b, 0x43, 0x3c, 0x1f, 0x68, 0xd2,
	0x79, 0x3f, 0x53, 0xfc, 0x34, 0x3a, 0x7f, 0x67, 0x49, 0x7f, 0x94, 0xbe, 0xb4, 0x68, 0xaf, 0x67,
	0xe0, 0x5e, 0xcf, 0xe0, 0xe3, 0x5e, 0x2f, 0x27, 0x5c, 0x6c, 0x1b, 0xb3, 0x6f, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x46, 0xe3, 0x38, 0xd9, 0x07, 0x00, 0x00,
}
