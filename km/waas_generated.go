package waaskm

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = make([]byte, 40)
var _NullReader = karmem.NewReader(_Null)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierMetadata    = 16498252775941464887
	PacketIdentifierSource      = 17594599419396936559
	PacketIdentifierDestination = 17000076967772502206
	PacketIdentifierInvocation  = 5708589682413831787
)

type Metadata struct {
	Key   string
	Value string
}

func NewMetadata() Metadata {
	return Metadata{}
}

func (x *Metadata) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierMetadata
}

func (x *Metadata) Reset() {
	x.Read((*MetadataViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Metadata) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Metadata) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__KeySize := uint(1 * len(x.Key))
	__KeyOffset, err := writer.Alloc(__KeySize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__KeyOffset))
	writer.Write4At(offset+0+4, uint32(__KeySize))
	writer.Write4At(offset+0+4+4, 1)
	__KeySlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Key)), __KeySize, __KeySize}
	writer.WriteAt(__KeyOffset, *(*[]byte)(unsafe.Pointer(&__KeySlice)))
	__ValueSize := uint(1 * len(x.Value))
	__ValueOffset, err := writer.Alloc(__ValueSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__ValueOffset))
	writer.Write4At(offset+12+4, uint32(__ValueSize))
	writer.Write4At(offset+12+4+4, 1)
	__ValueSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Value)), __ValueSize, __ValueSize}
	writer.WriteAt(__ValueOffset, *(*[]byte)(unsafe.Pointer(&__ValueSlice)))

	return offset, nil
}

func (x *Metadata) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewMetadataViewer(reader, 0), reader)
}

func (x *Metadata) Read(viewer *MetadataViewer, reader *karmem.Reader) {
	__KeyString := viewer.Key(reader)
	if x.Key != __KeyString {
		__KeyStringCopy := make([]byte, len(__KeyString))
		copy(__KeyStringCopy, __KeyString)
		x.Key = *(*string)(unsafe.Pointer(&__KeyStringCopy))
	}
	__ValueString := viewer.Value(reader)
	if x.Value != __ValueString {
		__ValueStringCopy := make([]byte, len(__ValueString))
		copy(__ValueStringCopy, __ValueString)
		x.Value = *(*string)(unsafe.Pointer(&__ValueStringCopy))
	}
}

type Source struct {
	Name     string
	Location string
}

func NewSource() Source {
	return Source{}
}

func (x *Source) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierSource
}

func (x *Source) Reset() {
	x.Read((*SourceViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Source) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Source) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(28))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	writer.Write4At(offset+4+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__LocationSize := uint(1 * len(x.Location))
	__LocationOffset, err := writer.Alloc(__LocationSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__LocationOffset))
	writer.Write4At(offset+16+4, uint32(__LocationSize))
	writer.Write4At(offset+16+4+4, 1)
	__LocationSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Location)), __LocationSize, __LocationSize}
	writer.WriteAt(__LocationOffset, *(*[]byte)(unsafe.Pointer(&__LocationSlice)))

	return offset, nil
}

func (x *Source) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewSourceViewer(reader, 0), reader)
}

func (x *Source) Read(viewer *SourceViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	__LocationString := viewer.Location(reader)
	if x.Location != __LocationString {
		__LocationStringCopy := make([]byte, len(__LocationString))
		copy(__LocationStringCopy, __LocationString)
		x.Location = *(*string)(unsafe.Pointer(&__LocationStringCopy))
	}
}

type Destination struct {
	Name     string
	Location string
}

func NewDestination() Destination {
	return Destination{}
}

func (x *Destination) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierDestination
}

func (x *Destination) Reset() {
	x.Read((*DestinationViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Destination) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Destination) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(32)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(28))
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__NameOffset))
	writer.Write4At(offset+4+4, uint32(__NameSize))
	writer.Write4At(offset+4+4+4, 1)
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__LocationSize := uint(1 * len(x.Location))
	__LocationOffset, err := writer.Alloc(__LocationSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__LocationOffset))
	writer.Write4At(offset+16+4, uint32(__LocationSize))
	writer.Write4At(offset+16+4+4, 1)
	__LocationSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Location)), __LocationSize, __LocationSize}
	writer.WriteAt(__LocationOffset, *(*[]byte)(unsafe.Pointer(&__LocationSlice)))

	return offset, nil
}

func (x *Destination) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewDestinationViewer(reader, 0), reader)
}

func (x *Destination) Read(viewer *DestinationViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	__LocationString := viewer.Location(reader)
	if x.Location != __LocationString {
		__LocationStringCopy := make([]byte, len(__LocationString))
		copy(__LocationStringCopy, __LocationString)
		x.Location = *(*string)(unsafe.Pointer(&__LocationStringCopy))
	}
}

type Invocation struct {
	Source      Source
	Destination Destination
	Payload     []byte
	Metadata    []Metadata
}

func NewInvocation() Invocation {
	return Invocation{}
}

func (x *Invocation) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierInvocation
}

func (x *Invocation) Reset() {
	x.Read((*InvocationViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Invocation) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Invocation) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(40)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	writer.Write4At(offset, uint32(36))
	__SourceSize := uint(32)
	__SourceOffset, err := writer.Alloc(__SourceSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+4, uint32(__SourceOffset))
	if _, err := x.Source.Write(writer, __SourceOffset); err != nil {
		return offset, err
	}
	__DestinationSize := uint(32)
	__DestinationOffset, err := writer.Alloc(__DestinationSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+8, uint32(__DestinationOffset))
	if _, err := x.Destination.Write(writer, __DestinationOffset); err != nil {
		return offset, err
	}
	__PayloadSize := uint(1 * len(x.Payload))
	__PayloadOffset, err := writer.Alloc(__PayloadSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+12, uint32(__PayloadOffset))
	writer.Write4At(offset+12+4, uint32(__PayloadSize))
	writer.Write4At(offset+12+4+4, 1)
	__PayloadSlice := *(*[3]uint)(unsafe.Pointer(&x.Payload))
	__PayloadSlice[1] = __PayloadSize
	__PayloadSlice[2] = __PayloadSize
	writer.WriteAt(__PayloadOffset, *(*[]byte)(unsafe.Pointer(&__PayloadSlice)))
	__MetadataSize := uint(32 * len(x.Metadata))
	__MetadataOffset, err := writer.Alloc(__MetadataSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+24, uint32(__MetadataOffset))
	writer.Write4At(offset+24+4, uint32(__MetadataSize))
	writer.Write4At(offset+24+4+4, 32)
	for i := range x.Metadata {
		if _, err := x.Metadata[i].Write(writer, __MetadataOffset); err != nil {
			return offset, err
		}
		__MetadataOffset += 32
	}

	return offset, nil
}

func (x *Invocation) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewInvocationViewer(reader, 0), reader)
}

func (x *Invocation) Read(viewer *InvocationViewer, reader *karmem.Reader) {
	x.Source.Read(viewer.Source(reader), reader)
	x.Destination.Read(viewer.Destination(reader), reader)
	__PayloadSlice := viewer.Payload(reader)
	__PayloadLen := len(__PayloadSlice)
	if __PayloadLen > cap(x.Payload) {
		x.Payload = append(x.Payload, make([]byte, __PayloadLen-len(x.Payload))...)
	}
	if __PayloadLen > len(x.Payload) {
		x.Payload = x.Payload[:__PayloadLen]
	}
	copy(x.Payload, __PayloadSlice)
	x.Payload = x.Payload[:__PayloadLen]
	__MetadataSlice := viewer.Metadata(reader)
	__MetadataLen := len(__MetadataSlice)
	if __MetadataLen > cap(x.Metadata) {
		x.Metadata = append(x.Metadata, make([]Metadata, __MetadataLen-len(x.Metadata))...)
	}
	if __MetadataLen > len(x.Metadata) {
		x.Metadata = x.Metadata[:__MetadataLen]
	}
	for i := 0; i < __MetadataLen; i++ {
		x.Metadata[i].Read(&__MetadataSlice[i], reader)
	}
	x.Metadata = x.Metadata[:__MetadataLen]
}

type MetadataViewer struct {
	_data [32]byte
}

func NewMetadataViewer(reader *karmem.Reader, offset uint32) (v *MetadataViewer) {
	if !reader.IsValidOffset(offset, 32) {
		return (*MetadataViewer)(unsafe.Pointer(&_Null))
	}
	v = (*MetadataViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *MetadataViewer) size() uint32 {
	return 32
}
func (x *MetadataViewer) Key(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *MetadataViewer) Value(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type SourceViewer struct {
	_data [32]byte
}

func NewSourceViewer(reader *karmem.Reader, offset uint32) (v *SourceViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*SourceViewer)(unsafe.Pointer(&_Null))
	}
	v = (*SourceViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*SourceViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *SourceViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *SourceViewer) Name(reader *karmem.Reader) (v string) {
	if 4+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *SourceViewer) Location(reader *karmem.Reader) (v string) {
	if 16+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type DestinationViewer struct {
	_data [32]byte
}

func NewDestinationViewer(reader *karmem.Reader, offset uint32) (v *DestinationViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*DestinationViewer)(unsafe.Pointer(&_Null))
	}
	v = (*DestinationViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*DestinationViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *DestinationViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *DestinationViewer) Name(reader *karmem.Reader) (v string) {
	if 4+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *DestinationViewer) Location(reader *karmem.Reader) (v string) {
	if 16+12 > x.size() {
		return v
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}

type InvocationViewer struct {
	_data [40]byte
}

func NewInvocationViewer(reader *karmem.Reader, offset uint32) (v *InvocationViewer) {
	if !reader.IsValidOffset(offset, 8) {
		return (*InvocationViewer)(unsafe.Pointer(&_Null))
	}
	v = (*InvocationViewer)(unsafe.Add(reader.Pointer, offset))
	if !reader.IsValidOffset(offset, v.size()) {
		return (*InvocationViewer)(unsafe.Pointer(&_Null))
	}
	return v
}

func (x *InvocationViewer) size() uint32 {
	return *(*uint32)(unsafe.Pointer(&x._data))
}
func (x *InvocationViewer) Source(reader *karmem.Reader) (v *SourceViewer) {
	if 4+4 > x.size() {
		return (*SourceViewer)(unsafe.Pointer(&_Null))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
	return NewSourceViewer(reader, offset)
}
func (x *InvocationViewer) Destination(reader *karmem.Reader) (v *DestinationViewer) {
	if 8+4 > x.size() {
		return (*DestinationViewer)(unsafe.Pointer(&_Null))
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 8))
	return NewDestinationViewer(reader, offset)
}
func (x *InvocationViewer) Payload(reader *karmem.Reader) (v []byte) {
	if 12+12 > x.size() {
		return []byte{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 12+4))
	if !reader.IsValidOffset(offset, size) {
		return []byte{}
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]byte)(unsafe.Pointer(&slice))
}
func (x *InvocationViewer) Metadata(reader *karmem.Reader) (v []MetadataViewer) {
	if 24+12 > x.size() {
		return []MetadataViewer{}
	}
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 24))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 24+4))
	if !reader.IsValidOffset(offset, size) {
		return []MetadataViewer{}
	}
	length := uintptr(size / 32)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*[]MetadataViewer)(unsafe.Pointer(&slice))
}
