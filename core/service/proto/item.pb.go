// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/item.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetItemsByLabelRequest struct {
	Label                string   `protobuf:"bytes,1,opt,name=Label,proto3" json:"Label,omitempty"`
	SortBy               string   `protobuf:"bytes,2,opt,name=SortBy,proto3" json:"SortBy,omitempty"`
	Begin                int64    `protobuf:"varint,3,opt,name=Begin,proto3" json:"Begin,omitempty"`
	End                  int64    `protobuf:"varint,4,opt,name=End,proto3" json:"End,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsByLabelRequest) Reset()         { *m = GetItemsByLabelRequest{} }
func (m *GetItemsByLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemsByLabelRequest) ProtoMessage()    {}
func (*GetItemsByLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{0}
}
func (m *GetItemsByLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsByLabelRequest.Unmarshal(m, b)
}
func (m *GetItemsByLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsByLabelRequest.Marshal(b, m, deterministic)
}
func (dst *GetItemsByLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsByLabelRequest.Merge(dst, src)
}
func (m *GetItemsByLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemsByLabelRequest.Size(m)
}
func (m *GetItemsByLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsByLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsByLabelRequest proto.InternalMessageInfo

func (m *GetItemsByLabelRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *GetItemsByLabelRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *GetItemsByLabelRequest) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *GetItemsByLabelRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type PagingShopGoodsResponse struct {
	Total                int64     `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Data                 []*SGoods `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PagingShopGoodsResponse) Reset()         { *m = PagingShopGoodsResponse{} }
func (m *PagingShopGoodsResponse) String() string { return proto.CompactTextString(m) }
func (*PagingShopGoodsResponse) ProtoMessage()    {}
func (*PagingShopGoodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{1}
}
func (m *PagingShopGoodsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingShopGoodsResponse.Unmarshal(m, b)
}
func (m *PagingShopGoodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingShopGoodsResponse.Marshal(b, m, deterministic)
}
func (dst *PagingShopGoodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingShopGoodsResponse.Merge(dst, src)
}
func (m *PagingShopGoodsResponse) XXX_Size() int {
	return xxx_messageInfo_PagingShopGoodsResponse.Size(m)
}
func (m *PagingShopGoodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingShopGoodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PagingShopGoodsResponse proto.InternalMessageInfo

func (m *PagingShopGoodsResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PagingShopGoodsResponse) GetData() []*SGoods {
	if m != nil {
		return m.Data
	}
	return nil
}

type PagingGoodsResponse struct {
	Total                int64               `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Data                 []*SUnifiedViewItem `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PagingGoodsResponse) Reset()         { *m = PagingGoodsResponse{} }
func (m *PagingGoodsResponse) String() string { return proto.CompactTextString(m) }
func (*PagingGoodsResponse) ProtoMessage()    {}
func (*PagingGoodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{2}
}
func (m *PagingGoodsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingGoodsResponse.Unmarshal(m, b)
}
func (m *PagingGoodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingGoodsResponse.Marshal(b, m, deterministic)
}
func (dst *PagingGoodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingGoodsResponse.Merge(dst, src)
}
func (m *PagingGoodsResponse) XXX_Size() int {
	return xxx_messageInfo_PagingGoodsResponse.Size(m)
}
func (m *PagingGoodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingGoodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PagingGoodsResponse proto.InternalMessageInfo

func (m *PagingGoodsResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PagingGoodsResponse) GetData() []*SUnifiedViewItem {
	if m != nil {
		return m.Data
	}
	return nil
}

// 获取商品请求
type GetItemsRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Begin      int64 `protobuf:"varint,2,opt,name=Begin,proto3" json:"Begin,omitempty"`
	End        int64 `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"`
	// 　是否随机
	Random               bool     `protobuf:"varint,4,opt,name=Random,proto3" json:"Random,omitempty"`
	Where                string   `protobuf:"bytes,5,opt,name=Where,proto3" json:"Where,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItemsRequest) Reset()         { *m = GetItemsRequest{} }
func (m *GetItemsRequest) String() string { return proto.CompactTextString(m) }
func (*GetItemsRequest) ProtoMessage()    {}
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{3}
}
func (m *GetItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItemsRequest.Unmarshal(m, b)
}
func (m *GetItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItemsRequest.Marshal(b, m, deterministic)
}
func (dst *GetItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItemsRequest.Merge(dst, src)
}
func (m *GetItemsRequest) XXX_Size() int {
	return xxx_messageInfo_GetItemsRequest.Size(m)
}
func (m *GetItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetItemsRequest proto.InternalMessageInfo

func (m *GetItemsRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *GetItemsRequest) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *GetItemsRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *GetItemsRequest) GetRandom() bool {
	if m != nil {
		return m.Random
	}
	return false
}

func (m *GetItemsRequest) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

// 简单的商品信息
type SGoods struct {
	ItemId     int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ProductId  int64 `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	VendorId   int64 `protobuf:"varint,3,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	ShopId     int64 `protobuf:"varint,4,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	CategoryId int32 `protobuf:"varint,5,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	// 标题
	Title      string `protobuf:"bytes,6,opt,name=Title,proto3" json:"Title,omitempty"`
	ShortTitle string `protobuf:"bytes,7,opt,name=ShortTitle,proto3" json:"ShortTitle,omitempty"`
	// 货号
	GoodsNo string `protobuf:"bytes,8,opt,name=GoodsNo,proto3" json:"GoodsNo,omitempty"`
	Image   string `protobuf:"bytes,9,opt,name=Image,proto3" json:"Image,omitempty"`
	// 定价
	RetailPrice float64 `protobuf:"fixed64,10,opt,name=RetailPrice,proto3" json:"RetailPrice,omitempty"`
	// 销售价
	Price float64 `protobuf:"fixed64,11,opt,name=Price,proto3" json:"Price,omitempty"`
	// 促销价
	PromPrice float64 `protobuf:"fixed64,12,opt,name=PromPrice,proto3" json:"PromPrice,omitempty"`
	// 价格区间
	PriceRange string `protobuf:"bytes,13,opt,name=PriceRange,proto3" json:"PriceRange,omitempty"`
	GoodsId    int64  `protobuf:"varint,14,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	SkuId      int64  `protobuf:"varint,15,opt,name=SkuId,proto3" json:"SkuId,omitempty"`
	IsPresent  bool   `protobuf:"varint,16,opt,name=IsPresent,proto3" json:"IsPresent,omitempty"`
	// 促销标志
	PromotionFlag int32 `protobuf:"varint,17,opt,name=PromotionFlag,proto3" json:"PromotionFlag,omitempty"`
	// 库存
	StockNum int32 `protobuf:"varint,18,opt,name=StockNum,proto3" json:"StockNum,omitempty"`
	// 已售件数
	SaleNum              int32    `protobuf:"varint,19,opt,name=SaleNum,proto3" json:"SaleNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SGoods) Reset()         { *m = SGoods{} }
func (m *SGoods) String() string { return proto.CompactTextString(m) }
func (*SGoods) ProtoMessage()    {}
func (*SGoods) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{4}
}
func (m *SGoods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SGoods.Unmarshal(m, b)
}
func (m *SGoods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SGoods.Marshal(b, m, deterministic)
}
func (dst *SGoods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SGoods.Merge(dst, src)
}
func (m *SGoods) XXX_Size() int {
	return xxx_messageInfo_SGoods.Size(m)
}
func (m *SGoods) XXX_DiscardUnknown() {
	xxx_messageInfo_SGoods.DiscardUnknown(m)
}

var xxx_messageInfo_SGoods proto.InternalMessageInfo

func (m *SGoods) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *SGoods) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SGoods) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SGoods) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SGoods) GetCategoryId() int32 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *SGoods) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SGoods) GetShortTitle() string {
	if m != nil {
		return m.ShortTitle
	}
	return ""
}

func (m *SGoods) GetGoodsNo() string {
	if m != nil {
		return m.GoodsNo
	}
	return ""
}

func (m *SGoods) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SGoods) GetRetailPrice() float64 {
	if m != nil {
		return m.RetailPrice
	}
	return 0
}

func (m *SGoods) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SGoods) GetPromPrice() float64 {
	if m != nil {
		return m.PromPrice
	}
	return 0
}

func (m *SGoods) GetPriceRange() string {
	if m != nil {
		return m.PriceRange
	}
	return ""
}

func (m *SGoods) GetGoodsId() int64 {
	if m != nil {
		return m.GoodsId
	}
	return 0
}

func (m *SGoods) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *SGoods) GetIsPresent() bool {
	if m != nil {
		return m.IsPresent
	}
	return false
}

func (m *SGoods) GetPromotionFlag() int32 {
	if m != nil {
		return m.PromotionFlag
	}
	return 0
}

func (m *SGoods) GetStockNum() int32 {
	if m != nil {
		return m.StockNum
	}
	return 0
}

func (m *SGoods) GetSaleNum() int32 {
	if m != nil {
		return m.SaleNum
	}
	return 0
}

// * SKU
type SSku struct {
	SkuId                int64    `protobuf:"zigzag64,1,opt,name=SkuId,proto3" json:"SkuId,omitempty"`
	ItemId               int64    `protobuf:"zigzag64,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ProductId            int64    `protobuf:"zigzag64,3,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	SpecData             string   `protobuf:"bytes,6,opt,name=SpecData,proto3" json:"SpecData,omitempty"`
	SpecWord             string   `protobuf:"bytes,7,opt,name=SpecWord,proto3" json:"SpecWord,omitempty"`
	Code                 string   `protobuf:"bytes,8,opt,name=Code,proto3" json:"Code,omitempty"`
	RetailPrice          float64  `protobuf:"fixed64,9,opt,name=RetailPrice,proto3" json:"RetailPrice,omitempty"`
	Price                float64  `protobuf:"fixed64,10,opt,name=Price,proto3" json:"Price,omitempty"`
	Cost                 float64  `protobuf:"fixed64,11,opt,name=Cost,proto3" json:"Cost,omitempty"`
	Weight               int32    `protobuf:"zigzag32,12,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Bulk                 int32    `protobuf:"zigzag32,13,opt,name=Bulk,proto3" json:"Bulk,omitempty"`
	Stock                int32    `protobuf:"zigzag32,14,opt,name=Stock,proto3" json:"Stock,omitempty"`
	SaleNum              int32    `protobuf:"zigzag32,15,opt,name=SaleNum,proto3" json:"SaleNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSku) Reset()         { *m = SSku{} }
func (m *SSku) String() string { return proto.CompactTextString(m) }
func (*SSku) ProtoMessage()    {}
func (*SSku) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{5}
}
func (m *SSku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSku.Unmarshal(m, b)
}
func (m *SSku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSku.Marshal(b, m, deterministic)
}
func (dst *SSku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSku.Merge(dst, src)
}
func (m *SSku) XXX_Size() int {
	return xxx_messageInfo_SSku.Size(m)
}
func (m *SSku) XXX_DiscardUnknown() {
	xxx_messageInfo_SSku.DiscardUnknown(m)
}

var xxx_messageInfo_SSku proto.InternalMessageInfo

func (m *SSku) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *SSku) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *SSku) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SSku) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SSku) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SSku) GetSpecData() string {
	if m != nil {
		return m.SpecData
	}
	return ""
}

func (m *SSku) GetSpecWord() string {
	if m != nil {
		return m.SpecWord
	}
	return ""
}

func (m *SSku) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SSku) GetRetailPrice() float64 {
	if m != nil {
		return m.RetailPrice
	}
	return 0
}

func (m *SSku) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SSku) GetCost() float64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *SSku) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *SSku) GetBulk() int32 {
	if m != nil {
		return m.Bulk
	}
	return 0
}

func (m *SSku) GetStock() int32 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *SSku) GetSaleNum() int32 {
	if m != nil {
		return m.SaleNum
	}
	return 0
}

// 统一的商品显示对象
type SUnifiedViewItem struct {
	ItemId               int64             `protobuf:"zigzag64,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ProductId            int64             `protobuf:"zigzag64,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	CatId                int32             `protobuf:"zigzag32,3,opt,name=CatId,proto3" json:"CatId,omitempty"`
	VendorId             int64             `protobuf:"zigzag64,4,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	BrandId              int32             `protobuf:"zigzag32,5,opt,name=BrandId,proto3" json:"BrandId,omitempty"`
	Title                string            `protobuf:"bytes,6,opt,name=Title,proto3" json:"Title,omitempty"`
	Code                 string            `protobuf:"bytes,7,opt,name=Code,proto3" json:"Code,omitempty"`
	Image                string            `protobuf:"bytes,8,opt,name=Image,proto3" json:"Image,omitempty"`
	Price                float64           `protobuf:"fixed64,9,opt,name=Price,proto3" json:"Price,omitempty"`
	PriceRange           string            `protobuf:"bytes,10,opt,name=PriceRange,proto3" json:"PriceRange,omitempty"`
	StockNum             int32             `protobuf:"zigzag32,11,opt,name=StockNum,proto3" json:"StockNum,omitempty"`
	ShelveState          int32             `protobuf:"zigzag32,12,opt,name=ShelveState,proto3" json:"ShelveState,omitempty"`
	ReviewState          int32             `protobuf:"zigzag32,13,opt,name=ReviewState,proto3" json:"ReviewState,omitempty"`
	UpdateTime           int64             `protobuf:"zigzag64,14,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	SkuArray             []*SSku           `protobuf:"bytes,15,rep,name=SkuArray,proto3" json:"SkuArray,omitempty"`
	Data                 map[string]string `protobuf:"bytes,16,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SUnifiedViewItem) Reset()         { *m = SUnifiedViewItem{} }
func (m *SUnifiedViewItem) String() string { return proto.CompactTextString(m) }
func (*SUnifiedViewItem) ProtoMessage()    {}
func (*SUnifiedViewItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{6}
}
func (m *SUnifiedViewItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SUnifiedViewItem.Unmarshal(m, b)
}
func (m *SUnifiedViewItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SUnifiedViewItem.Marshal(b, m, deterministic)
}
func (dst *SUnifiedViewItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SUnifiedViewItem.Merge(dst, src)
}
func (m *SUnifiedViewItem) XXX_Size() int {
	return xxx_messageInfo_SUnifiedViewItem.Size(m)
}
func (m *SUnifiedViewItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SUnifiedViewItem.DiscardUnknown(m)
}

var xxx_messageInfo_SUnifiedViewItem proto.InternalMessageInfo

func (m *SUnifiedViewItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *SUnifiedViewItem) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SUnifiedViewItem) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *SUnifiedViewItem) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SUnifiedViewItem) GetBrandId() int32 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *SUnifiedViewItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SUnifiedViewItem) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SUnifiedViewItem) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SUnifiedViewItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SUnifiedViewItem) GetPriceRange() string {
	if m != nil {
		return m.PriceRange
	}
	return ""
}

func (m *SUnifiedViewItem) GetStockNum() int32 {
	if m != nil {
		return m.StockNum
	}
	return 0
}

func (m *SUnifiedViewItem) GetShelveState() int32 {
	if m != nil {
		return m.ShelveState
	}
	return 0
}

func (m *SUnifiedViewItem) GetReviewState() int32 {
	if m != nil {
		return m.ReviewState
	}
	return 0
}

func (m *SUnifiedViewItem) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SUnifiedViewItem) GetSkuArray() []*SSku {
	if m != nil {
		return m.SkuArray
	}
	return nil
}

func (m *SUnifiedViewItem) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// 完整的商品信息
type SOldItem struct {
	ItemId               int64             `protobuf:"zigzag64,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ProductId            int64             `protobuf:"zigzag64,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	PromFlag             int32             `protobuf:"zigzag32,3,opt,name=PromFlag,proto3" json:"PromFlag,omitempty"`
	CatId                int32             `protobuf:"zigzag32,4,opt,name=CatId,proto3" json:"CatId,omitempty"`
	VendorId             int64             `protobuf:"zigzag64,5,opt,name=VendorId,proto3" json:"VendorId,omitempty"`
	BrandId              int32             `protobuf:"zigzag32,6,opt,name=BrandId,proto3" json:"BrandId,omitempty"`
	ShopId               int64             `protobuf:"zigzag64,7,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	ShopCatId            int32             `protobuf:"zigzag32,8,opt,name=ShopCatId,proto3" json:"ShopCatId,omitempty"`
	ExpressTid           int32             `protobuf:"zigzag32,9,opt,name=ExpressTid,proto3" json:"ExpressTid,omitempty"`
	Title                string            `protobuf:"bytes,10,opt,name=Title,proto3" json:"Title,omitempty"`
	ShortTitle           string            `protobuf:"bytes,11,opt,name=ShortTitle,proto3" json:"ShortTitle,omitempty"`
	Code                 string            `protobuf:"bytes,12,opt,name=Code,proto3" json:"Code,omitempty"`
	Image                string            `protobuf:"bytes,13,opt,name=Image,proto3" json:"Image,omitempty"`
	IsPresent            int32             `protobuf:"zigzag32,14,opt,name=IsPresent,proto3" json:"IsPresent,omitempty"`
	PriceRange           string            `protobuf:"bytes,15,opt,name=PriceRange,proto3" json:"PriceRange,omitempty"`
	StockNum             int32             `protobuf:"zigzag32,16,opt,name=StockNum,proto3" json:"StockNum,omitempty"`
	SaleNum              int32             `protobuf:"zigzag32,17,opt,name=SaleNum,proto3" json:"SaleNum,omitempty"`
	SkuNum               int32             `protobuf:"zigzag32,18,opt,name=SkuNum,proto3" json:"SkuNum,omitempty"`
	SkuId                int64             `protobuf:"zigzag64,19,opt,name=SkuId,proto3" json:"SkuId,omitempty"`
	Cost                 float64           `protobuf:"fixed64,20,opt,name=Cost,proto3" json:"Cost,omitempty"`
	Price                float64           `protobuf:"fixed64,21,opt,name=Price,proto3" json:"Price,omitempty"`
	RetailPrice          float64           `protobuf:"fixed64,22,opt,name=RetailPrice,proto3" json:"RetailPrice,omitempty"`
	Weight               int32             `protobuf:"zigzag32,23,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Bulk                 int32             `protobuf:"zigzag32,24,opt,name=Bulk,proto3" json:"Bulk,omitempty"`
	ShelveState          int32             `protobuf:"zigzag32,25,opt,name=ShelveState,proto3" json:"ShelveState,omitempty"`
	ReviewState          int32             `protobuf:"zigzag32,26,opt,name=ReviewState,proto3" json:"ReviewState,omitempty"`
	ReviewRemark         string            `protobuf:"bytes,27,opt,name=ReviewRemark,proto3" json:"ReviewRemark,omitempty"`
	SortNum              int32             `protobuf:"zigzag32,28,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	CreateTime           int64             `protobuf:"zigzag64,29,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime           int64             `protobuf:"zigzag64,30,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	PromPrice            float64           `protobuf:"fixed64,31,opt,name=PromPrice,proto3" json:"PromPrice,omitempty"`
	SkuArray             []*SSku           `protobuf:"bytes,32,rep,name=SkuArray,proto3" json:"SkuArray,omitempty"`
	Data                 map[string]string `protobuf:"bytes,33,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SOldItem) Reset()         { *m = SOldItem{} }
func (m *SOldItem) String() string { return proto.CompactTextString(m) }
func (*SOldItem) ProtoMessage()    {}
func (*SOldItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{7}
}
func (m *SOldItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SOldItem.Unmarshal(m, b)
}
func (m *SOldItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SOldItem.Marshal(b, m, deterministic)
}
func (dst *SOldItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SOldItem.Merge(dst, src)
}
func (m *SOldItem) XXX_Size() int {
	return xxx_messageInfo_SOldItem.Size(m)
}
func (m *SOldItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SOldItem.DiscardUnknown(m)
}

var xxx_messageInfo_SOldItem proto.InternalMessageInfo

func (m *SOldItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *SOldItem) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *SOldItem) GetPromFlag() int32 {
	if m != nil {
		return m.PromFlag
	}
	return 0
}

func (m *SOldItem) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *SOldItem) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *SOldItem) GetBrandId() int32 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *SOldItem) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *SOldItem) GetShopCatId() int32 {
	if m != nil {
		return m.ShopCatId
	}
	return 0
}

func (m *SOldItem) GetExpressTid() int32 {
	if m != nil {
		return m.ExpressTid
	}
	return 0
}

func (m *SOldItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SOldItem) GetShortTitle() string {
	if m != nil {
		return m.ShortTitle
	}
	return ""
}

func (m *SOldItem) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SOldItem) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SOldItem) GetIsPresent() int32 {
	if m != nil {
		return m.IsPresent
	}
	return 0
}

func (m *SOldItem) GetPriceRange() string {
	if m != nil {
		return m.PriceRange
	}
	return ""
}

func (m *SOldItem) GetStockNum() int32 {
	if m != nil {
		return m.StockNum
	}
	return 0
}

func (m *SOldItem) GetSaleNum() int32 {
	if m != nil {
		return m.SaleNum
	}
	return 0
}

func (m *SOldItem) GetSkuNum() int32 {
	if m != nil {
		return m.SkuNum
	}
	return 0
}

func (m *SOldItem) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *SOldItem) GetCost() float64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *SOldItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SOldItem) GetRetailPrice() float64 {
	if m != nil {
		return m.RetailPrice
	}
	return 0
}

func (m *SOldItem) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *SOldItem) GetBulk() int32 {
	if m != nil {
		return m.Bulk
	}
	return 0
}

func (m *SOldItem) GetShelveState() int32 {
	if m != nil {
		return m.ShelveState
	}
	return 0
}

func (m *SOldItem) GetReviewState() int32 {
	if m != nil {
		return m.ReviewState
	}
	return 0
}

func (m *SOldItem) GetReviewRemark() string {
	if m != nil {
		return m.ReviewRemark
	}
	return ""
}

func (m *SOldItem) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SOldItem) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SOldItem) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SOldItem) GetPromPrice() float64 {
	if m != nil {
		return m.PromPrice
	}
	return 0
}

func (m *SOldItem) GetSkuArray() []*SSku {
	if m != nil {
		return m.SkuArray
	}
	return nil
}

func (m *SOldItem) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type ItemLabelListResponse struct {
	Value                []*SItemLabel `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ItemLabelListResponse) Reset()         { *m = ItemLabelListResponse{} }
func (m *ItemLabelListResponse) String() string { return proto.CompactTextString(m) }
func (*ItemLabelListResponse) ProtoMessage()    {}
func (*ItemLabelListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{8}
}
func (m *ItemLabelListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemLabelListResponse.Unmarshal(m, b)
}
func (m *ItemLabelListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemLabelListResponse.Marshal(b, m, deterministic)
}
func (dst *ItemLabelListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemLabelListResponse.Merge(dst, src)
}
func (m *ItemLabelListResponse) XXX_Size() int {
	return xxx_messageInfo_ItemLabelListResponse.Size(m)
}
func (m *ItemLabelListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemLabelListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemLabelListResponse proto.InternalMessageInfo

func (m *ItemLabelListResponse) GetValue() []*SItemLabel {
	if m != nil {
		return m.Value
	}
	return nil
}

// 销售标签
type SItemLabel struct {
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 标签名
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 标签代码
	TagCode string `protobuf:"bytes,5,opt,name=TagCode,proto3" json:"TagCode,omitempty"`
	// 商品的遮盖图
	LabelImage string `protobuf:"bytes,3,opt,name=LabelImage,proto3" json:"LabelImage,omitempty"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,4,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SItemLabel) Reset()         { *m = SItemLabel{} }
func (m *SItemLabel) String() string { return proto.CompactTextString(m) }
func (*SItemLabel) ProtoMessage()    {}
func (*SItemLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_60a8a7a280474021, []int{9}
}
func (m *SItemLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SItemLabel.Unmarshal(m, b)
}
func (m *SItemLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SItemLabel.Marshal(b, m, deterministic)
}
func (dst *SItemLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SItemLabel.Merge(dst, src)
}
func (m *SItemLabel) XXX_Size() int {
	return xxx_messageInfo_SItemLabel.Size(m)
}
func (m *SItemLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_SItemLabel.DiscardUnknown(m)
}

var xxx_messageInfo_SItemLabel proto.InternalMessageInfo

func (m *SItemLabel) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SItemLabel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SItemLabel) GetTagCode() string {
	if m != nil {
		return m.TagCode
	}
	return ""
}

func (m *SItemLabel) GetLabelImage() string {
	if m != nil {
		return m.LabelImage
	}
	return ""
}

func (m *SItemLabel) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*GetItemsByLabelRequest)(nil), "GetItemsByLabelRequest")
	proto.RegisterType((*PagingShopGoodsResponse)(nil), "PagingShopGoodsResponse")
	proto.RegisterType((*PagingGoodsResponse)(nil), "PagingGoodsResponse")
	proto.RegisterType((*GetItemsRequest)(nil), "GetItemsRequest")
	proto.RegisterType((*SGoods)(nil), "SGoods")
	proto.RegisterType((*SSku)(nil), "SSku")
	proto.RegisterType((*SUnifiedViewItem)(nil), "SUnifiedViewItem")
	proto.RegisterMapType((map[string]string)(nil), "SUnifiedViewItem.DataEntry")
	proto.RegisterType((*SOldItem)(nil), "SOldItem")
	proto.RegisterMapType((map[string]string)(nil), "SOldItem.DataEntry")
	proto.RegisterType((*ItemLabelListResponse)(nil), "ItemLabelListResponse")
	proto.RegisterType((*SItemLabel)(nil), "SItemLabel")
}

func init() { proto.RegisterFile("message/item.proto", fileDescriptor_item_60a8a7a280474021) }

var fileDescriptor_item_60a8a7a280474021 = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0xfe, 0x9a, 0x64, 0xd2, 0x6e, 0x1b, 0x77, 0xb7, 0x3b, 0xb4, 0x65, 0x49, 0x23, 0x10,
	0xbd, 0xca, 0x4a, 0x70, 0x01, 0x5a, 0xae, 0x48, 0x29, 0xab, 0x48, 0x55, 0xa9, 0xc6, 0xdd, 0x56,
	0xe2, 0x6e, 0x5a, 0x1f, 0x1c, 0x2b, 0xb6, 0x27, 0xd8, 0x93, 0x2e, 0xb9, 0x47, 0xe2, 0x11, 0xe0,
	0xe5, 0x78, 0x09, 0x9e, 0x00, 0xcd, 0x99, 0x19, 0x7b, 0xec, 0xb4, 0x0b, 0x12, 0x5c, 0x65, 0xbe,
	0x73, 0x32, 0x7f, 0xdf, 0xf9, 0xbe, 0x33, 0x26, 0x5e, 0x02, 0x79, 0xce, 0x43, 0x78, 0x1d, 0x49,
	0x48, 0x26, 0xcb, 0x4c, 0x48, 0x31, 0x4e, 0xc9, 0xc1, 0x5b, 0x90, 0x33, 0x09, 0x49, 0x3e, 0x5d,
	0x5f, 0xf0, 0x3b, 0x88, 0x19, 0xfc, 0xbc, 0x82, 0x5c, 0x7a, 0xcf, 0x49, 0x07, 0x31, 0x6d, 0x8c,
	0x1a, 0xa7, 0x7d, 0xa6, 0x81, 0x77, 0x40, 0xb6, 0x7c, 0x91, 0xc9, 0xe9, 0x9a, 0x36, 0x31, 0x6c,
	0x90, 0xfa, 0xf7, 0x14, 0xc2, 0x28, 0xa5, 0xad, 0x51, 0xe3, 0xb4, 0xc5, 0x34, 0xf0, 0xf6, 0x48,
	0xeb, 0x3c, 0x0d, 0x68, 0x1b, 0x63, 0x6a, 0x38, 0xbe, 0x20, 0x2f, 0xaf, 0x78, 0x18, 0xa5, 0xa1,
	0x3f, 0x17, 0xcb, 0xb7, 0x42, 0x04, 0x39, 0x83, 0x7c, 0x29, 0xd2, 0x1c, 0xd4, 0x12, 0xd7, 0x42,
	0x72, 0xbd, 0x61, 0x8b, 0x69, 0xe0, 0x1d, 0x91, 0xf6, 0x77, 0x5c, 0x72, 0xda, 0x1c, 0xb5, 0x4e,
	0x07, 0x5f, 0x74, 0x27, 0xbe, 0x9e, 0x84, 0xc1, 0x31, 0x23, 0xfb, 0x7a, 0xb5, 0x7f, 0xb3, 0xd2,
	0x67, 0x95, 0x95, 0x86, 0x13, 0xff, 0x5d, 0x1a, 0xfd, 0x14, 0x41, 0x70, 0x13, 0xc1, 0x7b, 0x45,
	0x80, 0x59, 0xf3, 0xb7, 0x06, 0xd9, 0xb5, 0x94, 0x58, 0x2e, 0x5e, 0x11, 0x72, 0xc6, 0x25, 0x84,
	0x22, 0x5b, 0xcf, 0x02, 0xb3, 0xaa, 0x13, 0x29, 0x6f, 0xdf, 0x7c, 0xe4, 0xf6, 0xad, 0xe2, 0xf6,
	0x8a, 0x3d, 0xc6, 0xd3, 0x40, 0x24, 0x48, 0x49, 0x8f, 0x19, 0xa4, 0xe6, 0xdf, 0xce, 0x21, 0x03,
	0xda, 0xd1, 0x5c, 0x23, 0x18, 0xff, 0xd1, 0x26, 0x5b, 0xfa, 0xba, 0x6a, 0xa2, 0x3a, 0x50, 0xb1,
	0xb9, 0x41, 0xde, 0x31, 0xe9, 0x5f, 0x65, 0x22, 0x58, 0xdd, 0xcb, 0x59, 0x60, 0x36, 0x2f, 0x03,
	0xde, 0x21, 0xe9, 0xdd, 0x40, 0x1a, 0x88, 0x6c, 0x66, 0x4f, 0x51, 0x60, 0x2c, 0xe4, 0x5c, 0x2c,
	0x67, 0xb6, 0x3a, 0x06, 0xd5, 0xae, 0xaa, 0xce, 0xd3, 0xa9, 0x5f, 0xf5, 0x3a, 0x92, 0x31, 0xd0,
	0x2d, 0x7d, 0x54, 0x04, 0x6a, 0x96, 0x3f, 0x17, 0x99, 0xd4, 0xa9, 0x2e, 0xa6, 0x9c, 0x88, 0x47,
	0x49, 0x17, 0x2f, 0x72, 0x29, 0x68, 0x0f, 0x93, 0x16, 0xaa, 0xf5, 0x66, 0x09, 0x0f, 0x81, 0xf6,
	0xf5, 0x7a, 0x08, 0xbc, 0x11, 0x19, 0x30, 0x90, 0x3c, 0x8a, 0xaf, 0xb2, 0xe8, 0x1e, 0x28, 0x19,
	0x35, 0x4e, 0x1b, 0xcc, 0x0d, 0xa9, 0x79, 0x3a, 0x37, 0xc0, 0x9c, 0x06, 0x86, 0x8f, 0x44, 0x67,
	0xb6, 0x31, 0x53, 0x06, 0xd4, 0x29, 0x71, 0xc0, 0x78, 0x1a, 0x02, 0xdd, 0xd1, 0xa7, 0x2c, 0x23,
	0xc5, 0x29, 0x67, 0x01, 0x7d, 0x86, 0xa4, 0x58, 0xa8, 0x76, 0xf3, 0x17, 0xab, 0x59, 0x40, 0x77,
	0x75, 0x81, 0x11, 0xa8, 0xdd, 0x66, 0xf9, 0x55, 0x06, 0x39, 0xa4, 0x92, 0xee, 0x61, 0x45, 0xcb,
	0x80, 0xf7, 0x29, 0xd9, 0x51, 0x5b, 0x0b, 0x19, 0x89, 0xf4, 0xfb, 0x98, 0x87, 0x74, 0x88, 0x64,
	0x56, 0x83, 0xaa, 0x46, 0xbe, 0x14, 0xf7, 0x8b, 0xcb, 0x55, 0x42, 0x3d, 0xfc, 0x43, 0x81, 0xd5,
	0x79, 0x7c, 0x1e, 0x83, 0x4a, 0xed, 0x63, 0xca, 0xc2, 0xf1, 0x5f, 0x4d, 0xd2, 0xf6, 0xfd, 0xc5,
	0xaa, 0x3c, 0x98, 0xd2, 0x85, 0x67, 0x0f, 0x56, 0xca, 0xa5, 0x89, 0xe1, 0x47, 0xe5, 0xd2, 0xc2,
	0x94, 0x23, 0x97, 0xa2, 0xb4, 0x6d, 0xb7, 0xb4, 0x45, 0x81, 0x3a, 0x6e, 0x81, 0xd4, 0xb1, 0x97,
	0x70, 0x8f, 0x86, 0xd2, 0x4a, 0x28, 0xb0, 0xcd, 0xdd, 0x8a, 0x2c, 0x30, 0x52, 0x28, 0xb0, 0xe7,
	0x91, 0xf6, 0x99, 0x08, 0xc0, 0xa8, 0x00, 0xc7, 0xf5, 0x62, 0xf7, 0x3f, 0x50, 0x6c, 0xe2, 0x16,
	0x1b, 0xd7, 0xca, 0xa5, 0x51, 0x00, 0x8e, 0xd5, 0xcd, 0x6f, 0x21, 0x0a, 0xe7, 0x12, 0xab, 0x3f,
	0x64, 0x06, 0xa9, 0xff, 0x4e, 0x57, 0xf1, 0x02, 0x8b, 0x3e, 0x64, 0x38, 0x46, 0xee, 0x14, 0xd5,
	0x58, 0xec, 0x21, 0xd3, 0xc0, 0x25, 0x7d, 0x17, 0xe3, 0x05, 0xe9, 0xbf, 0xb7, 0xc9, 0x5e, 0xbd,
	0x69, 0xd4, 0x9c, 0xe9, 0x3d, 0xed, 0xcc, 0x3a, 0xd5, 0x67, 0xdc, 0x16, 0x61, 0xc8, 0x34, 0xa8,
	0xf8, 0xb5, 0x8d, 0x53, 0x4a, 0xbf, 0x52, 0xd2, 0x9d, 0x66, 0x3c, 0x0d, 0x8c, 0x29, 0x87, 0xcc,
	0xc2, 0x27, 0x1c, 0x69, 0x89, 0xee, 0x3a, 0x44, 0x17, 0xa5, 0xec, 0xb9, 0xa5, 0x2c, 0xc8, 0xed,
	0xbb, 0xe4, 0x56, 0xbd, 0x42, 0x36, 0xbc, 0xe2, 0xea, 0x76, 0x80, 0x07, 0x2a, 0x75, 0x3b, 0x22,
	0x03, 0x7f, 0x0e, 0xf1, 0x03, 0xf8, 0x92, 0x4b, 0x30, 0x95, 0x70, 0x43, 0xba, 0xe4, 0x0f, 0x11,
	0xbc, 0xd7, 0xff, 0xd0, 0x55, 0x71, 0x43, 0x6a, 0xff, 0x77, 0xcb, 0x80, 0x4b, 0xb8, 0x8e, 0x12,
	0xc0, 0x0a, 0x79, 0xcc, 0x89, 0x78, 0x27, 0xa4, 0xe7, 0x2f, 0x56, 0xdf, 0x66, 0x19, 0x5f, 0xd3,
	0x5d, 0xec, 0xe8, 0x9d, 0x89, 0x72, 0x04, 0x2b, 0xc2, 0xde, 0x6b, 0xd3, 0xf0, 0xf7, 0x30, 0x7d,
	0xb4, 0xd1, 0xf0, 0x27, 0x2a, 0x7b, 0x9e, 0xca, 0x6c, 0xad, 0x5b, 0xff, 0xe1, 0x57, 0xa4, 0x5f,
	0x84, 0x54, 0xf7, 0x5e, 0xc0, 0xda, 0xbc, 0x7e, 0x6a, 0xa8, 0x88, 0x7a, 0xe0, 0xf1, 0x0a, 0xcc,
	0xd3, 0xa7, 0xc1, 0x9b, 0xe6, 0xd7, 0x8d, 0xf1, 0x9f, 0x5d, 0xd2, 0xf3, 0x7f, 0x88, 0x83, 0xff,
	0xa0, 0x88, 0x43, 0xd2, 0x53, 0x8d, 0x01, 0x1b, 0x85, 0x16, 0x45, 0x81, 0x4b, 0xb5, 0xb4, 0x9f,
	0x52, 0x4b, 0xe7, 0x69, 0xb5, 0x6c, 0x55, 0xd5, 0x52, 0xf6, 0xfd, 0xae, 0x3e, 0x9d, 0xe9, 0xfb,
	0xc7, 0xa4, 0xaf, 0x46, 0x7a, 0x9f, 0x1e, 0xce, 0x29, 0x03, 0xaa, 0x1a, 0xe7, 0xbf, 0x2c, 0x33,
	0xc8, 0xf3, 0xeb, 0x28, 0x40, 0xa1, 0x0c, 0x99, 0x13, 0x29, 0x35, 0x48, 0x9e, 0x7e, 0x15, 0x06,
	0x1b, 0xaf, 0x82, 0xd5, 0xe8, 0xf6, 0x63, 0x1a, 0xdd, 0x71, 0x35, 0x5a, 0xe9, 0xb4, 0xda, 0xae,
	0x4e, 0xa7, 0xad, 0x6a, 0x75, 0xf7, 0x83, 0x5a, 0xdd, 0xab, 0x69, 0xd5, 0xb1, 0xfb, 0xb0, 0x62,
	0x77, 0x64, 0x6a, 0xb1, 0xb2, 0x7d, 0x79, 0xc8, 0x0c, 0x2a, 0x5b, 0xee, 0xbe, 0xdb, 0x72, 0x6d,
	0x33, 0x7a, 0xee, 0x34, 0xa3, 0xc2, 0x59, 0x2f, 0x5c, 0x67, 0xd5, 0xda, 0xdd, 0xc1, 0x66, 0xbb,
	0x2b, 0x9b, 0xd8, 0xcb, 0x47, 0x9b, 0x18, 0x75, 0x9a, 0x58, 0xcd, 0x6b, 0x1f, 0xfd, 0xa3, 0xd7,
	0x0e, 0x37, 0xbd, 0x36, 0x26, 0xdb, 0x1a, 0x32, 0x48, 0x78, 0xb6, 0xa0, 0x47, 0xc8, 0x60, 0x25,
	0x86, 0x3c, 0x89, 0x4c, 0x2a, 0x3a, 0x8e, 0x0d, 0x4f, 0x1a, 0xe2, 0x17, 0x43, 0x06, 0xd6, 0xa9,
	0x1f, 0x6b, 0xa7, 0x96, 0x91, 0x9a, 0x93, 0x5f, 0x6d, 0x38, 0xb9, 0xf2, 0x66, 0x7f, 0x52, 0x7f,
	0xb3, 0x5d, 0x9f, 0x8f, 0x1e, 0xf7, 0xf9, 0xe7, 0xc6, 0xe7, 0x27, 0x98, 0xde, 0x9f, 0x58, 0x27,
	0xfe, 0x7f, 0xfe, 0x7e, 0x43, 0x5e, 0xa8, 0x05, 0xf1, 0x13, 0xf8, 0x22, 0xca, 0x65, 0xf1, 0xa5,
	0x79, 0x42, 0x3a, 0x37, 0x38, 0xa5, 0x81, 0x7b, 0x0f, 0x26, 0x7e, 0xf1, 0x3f, 0xa6, 0x33, 0xe3,
	0x5f, 0x1b, 0x84, 0x94, 0x51, 0xef, 0x19, 0x69, 0x9a, 0xce, 0xd0, 0x61, 0x4d, 0xad, 0x9b, 0x4b,
	0x9e, 0xd8, 0x3d, 0x71, 0xac, 0xb8, 0xbe, 0xe6, 0x21, 0x5a, 0x43, 0x3f, 0xba, 0x16, 0x2a, 0x2e,
	0x71, 0x19, 0x6d, 0x91, 0x96, 0x56, 0x7a, 0x19, 0x51, 0x33, 0xcf, 0x53, 0x7e, 0x17, 0x43, 0x60,
	0xbe, 0x30, 0x2d, 0x9c, 0xf6, 0x7f, 0xec, 0x4e, 0xbe, 0xc1, 0x6f, 0xfe, 0xbb, 0x2d, 0xfc, 0xf9,
	0xf2, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x55, 0xa8, 0x3b, 0x10, 0x0c, 0x00, 0x00,
}