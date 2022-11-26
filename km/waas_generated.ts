
import * as karmem from 'karmem/assemblyscript/karmem'

let _Null = new StaticArray<u8>(40)
let _NullReader = karmem.NewReader(_Null)
export type PacketIdentifier = u64
export const PacketIdentifierMetadata = 16498252775941464887
export const PacketIdentifierSource = 17594599419396936559
export const PacketIdentifierDestination = 17000076967772502206
export const PacketIdentifierInvocation = 5708589682413831787



export class Metadata {
    Key: string;
    Value: string;

    static PacketIdentifier() : PacketIdentifier {
        return PacketIdentifierMetadata
    }

    static Reset(x: Metadata): void {
        this.Read(x, changetype<MetadataViewer>(changetype<usize>(_Null)), _NullReader)
    }

    @inline
    static WriteAsRoot(x: Metadata, writer: karmem.Writer): void {
        this.Write(x, writer, 0);
    }

    static Write(x: Metadata, writer: karmem.Writer, start: u32): boolean {
        let offset = start;
        const size: u32 = 32;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == 0xFFFFFFFF) {
                return false;
            }
        }
        const __KeyString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Key, false))
        const __KeySize: u32 = 1 * __KeyString.length;
        let __KeyOffset = writer.Alloc(__KeySize);
        if (__KeyOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +0, __KeyOffset);
        writer.WriteAt<u32>(offset +0 +4, __KeySize);
        writer.WriteAt<u32>(offset + 0 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__KeyOffset, __KeyString);
        const __ValueString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Value, false))
        const __ValueSize: u32 = 1 * __ValueString.length;
        let __ValueOffset = writer.Alloc(__ValueSize);
        if (__ValueOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +12, __ValueOffset);
        writer.WriteAt<u32>(offset +12 +4, __ValueSize);
        writer.WriteAt<u32>(offset + 12 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__ValueOffset, __ValueString);

        return true
    }

    @inline
    static ReadAsRoot(x: Metadata, reader: karmem.Reader) : void {
        this.Read(x, NewMetadataViewer(reader, 0), reader);
    }

    @inline
    static Read(x: Metadata, viewer: MetadataViewer, reader: karmem.Reader) : void {
    x.Key = viewer.Key(reader);
    x.Value = viewer.Value(reader);
    }

}

export function NewMetadata(): Metadata {
    let x: Metadata = {
    Key: "",
    Value: "",
    }
    return x;
}

export class Source {
    Name: string;
    Location: string;

    static PacketIdentifier() : PacketIdentifier {
        return PacketIdentifierSource
    }

    static Reset(x: Source): void {
        this.Read(x, changetype<SourceViewer>(changetype<usize>(_Null)), _NullReader)
    }

    @inline
    static WriteAsRoot(x: Source, writer: karmem.Writer): void {
        this.Write(x, writer, 0);
    }

    static Write(x: Source, writer: karmem.Writer, start: u32): boolean {
        let offset = start;
        const size: u32 = 32;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == 0xFFFFFFFF) {
                return false;
            }
        }
        writer.WriteAt<u32>(offset, 28);
        const __NameString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Name, false))
        const __NameSize: u32 = 1 * __NameString.length;
        let __NameOffset = writer.Alloc(__NameSize);
        if (__NameOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +4, __NameOffset);
        writer.WriteAt<u32>(offset +4 +4, __NameSize);
        writer.WriteAt<u32>(offset + 4 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__NameOffset, __NameString);
        const __LocationString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Location, false))
        const __LocationSize: u32 = 1 * __LocationString.length;
        let __LocationOffset = writer.Alloc(__LocationSize);
        if (__LocationOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +16, __LocationOffset);
        writer.WriteAt<u32>(offset +16 +4, __LocationSize);
        writer.WriteAt<u32>(offset + 16 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__LocationOffset, __LocationString);

        return true
    }

    @inline
    static ReadAsRoot(x: Source, reader: karmem.Reader) : void {
        this.Read(x, NewSourceViewer(reader, 0), reader);
    }

    @inline
    static Read(x: Source, viewer: SourceViewer, reader: karmem.Reader) : void {
    x.Name = viewer.Name(reader);
    x.Location = viewer.Location(reader);
    }

}

export function NewSource(): Source {
    let x: Source = {
    Name: "",
    Location: "",
    }
    return x;
}

export class Destination {
    Name: string;
    Location: string;

    static PacketIdentifier() : PacketIdentifier {
        return PacketIdentifierDestination
    }

    static Reset(x: Destination): void {
        this.Read(x, changetype<DestinationViewer>(changetype<usize>(_Null)), _NullReader)
    }

    @inline
    static WriteAsRoot(x: Destination, writer: karmem.Writer): void {
        this.Write(x, writer, 0);
    }

    static Write(x: Destination, writer: karmem.Writer, start: u32): boolean {
        let offset = start;
        const size: u32 = 32;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == 0xFFFFFFFF) {
                return false;
            }
        }
        writer.WriteAt<u32>(offset, 28);
        const __NameString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Name, false))
        const __NameSize: u32 = 1 * __NameString.length;
        let __NameOffset = writer.Alloc(__NameSize);
        if (__NameOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +4, __NameOffset);
        writer.WriteAt<u32>(offset +4 +4, __NameSize);
        writer.WriteAt<u32>(offset + 4 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__NameOffset, __NameString);
        const __LocationString : Uint8Array = Uint8Array.wrap(String.UTF8.encode(x.Location, false))
        const __LocationSize: u32 = 1 * __LocationString.length;
        let __LocationOffset = writer.Alloc(__LocationSize);
        if (__LocationOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +16, __LocationOffset);
        writer.WriteAt<u32>(offset +16 +4, __LocationSize);
        writer.WriteAt<u32>(offset + 16 + 4 + 4, 1)
        writer.WriteSliceAt<Uint8Array>(__LocationOffset, __LocationString);

        return true
    }

    @inline
    static ReadAsRoot(x: Destination, reader: karmem.Reader) : void {
        this.Read(x, NewDestinationViewer(reader, 0), reader);
    }

    @inline
    static Read(x: Destination, viewer: DestinationViewer, reader: karmem.Reader) : void {
    x.Name = viewer.Name(reader);
    x.Location = viewer.Location(reader);
    }

}

export function NewDestination(): Destination {
    let x: Destination = {
    Name: "",
    Location: "",
    }
    return x;
}

export class Invocation {
    Source: Source;
    Destination: Destination;
    Payload: Array<u8>;
    Metadata: Array<Metadata>;

    static PacketIdentifier() : PacketIdentifier {
        return PacketIdentifierInvocation
    }

    static Reset(x: Invocation): void {
        this.Read(x, changetype<InvocationViewer>(changetype<usize>(_Null)), _NullReader)
    }

    @inline
    static WriteAsRoot(x: Invocation, writer: karmem.Writer): void {
        this.Write(x, writer, 0);
    }

    static Write(x: Invocation, writer: karmem.Writer, start: u32): boolean {
        let offset = start;
        const size: u32 = 40;
        if (offset == 0) {
            offset = writer.Alloc(size);
            if (offset == 0xFFFFFFFF) {
                return false;
            }
        }
        writer.WriteAt<u32>(offset, 36);
        const __SourceSize: u32 = 32;
        let __SourceOffset = writer.Alloc(__SourceSize);
        if (__SourceOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +4, __SourceOffset);
        if (!Source.Write(x.Source, writer, __SourceOffset)) {
            return false;
        }
        const __DestinationSize: u32 = 32;
        let __DestinationOffset = writer.Alloc(__DestinationSize);
        if (__DestinationOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +8, __DestinationOffset);
        if (!Destination.Write(x.Destination, writer, __DestinationOffset)) {
            return false;
        }
        const __PayloadSize: u32 = 1 * x.Payload.length;
        let __PayloadOffset = writer.Alloc(__PayloadSize);
        if (__PayloadOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +12, __PayloadOffset);
        writer.WriteAt<u32>(offset +12 +4, __PayloadSize);
        writer.WriteAt<u32>(offset + 12 + 4 + 4, 1)
        writer.WriteSliceAt<Array<u8>>(__PayloadOffset, x.Payload);
        const __MetadataSize: u32 = 32 * x.Metadata.length;
        let __MetadataOffset = writer.Alloc(__MetadataSize);
        if (__MetadataOffset == 0) {
            return false;
        }
        writer.WriteAt<u32>(offset +24, __MetadataOffset);
        writer.WriteAt<u32>(offset +24 +4, __MetadataSize);
        writer.WriteAt<u32>(offset + 24 + 4 + 4, 32)
        let __MetadataLen = x.Metadata.length;
        for (let i = 0; i < __MetadataLen; i++) {
            if (!Metadata.Write(x.Metadata[i], writer, __MetadataOffset)) {
                return false;
            }
            __MetadataOffset += 32;
        }

        return true
    }

    @inline
    static ReadAsRoot(x: Invocation, reader: karmem.Reader) : void {
        this.Read(x, NewInvocationViewer(reader, 0), reader);
    }

    @inline
    static Read(x: Invocation, viewer: InvocationViewer, reader: karmem.Reader) : void {
    Source.Read(x.Source, viewer.Source(reader), reader);
    Destination.Read(x.Destination, viewer.Destination(reader), reader);
    let __PayloadSlice = viewer.Payload(reader);
    let __PayloadLen = __PayloadSlice.length;
    let __PayloadDestLen = x.Payload.length;
    if (__PayloadLen > __PayloadDestLen) {
        x.Payload.length = __PayloadDestLen
        for (let i = __PayloadDestLen; i < __PayloadLen; i++) {
            x.Payload.push(0);
        }
    }
    for (let i = 0; i < __PayloadLen; i++) {
        x.Payload[i] = __PayloadSlice[i];
    }
    x.Payload.length = __PayloadLen;
    let __MetadataSlice = viewer.Metadata(reader);
    let __MetadataLen = __MetadataSlice.length;
    let __MetadataDestLen = x.Metadata.length;
    if (__MetadataLen > __MetadataDestLen) {
        x.Metadata.length = __MetadataDestLen
        for (let i = __MetadataDestLen; i < __MetadataLen; i++) {
            x.Metadata.push(NewMetadata());
        }
    }
    for (let i = 0; i < __MetadataLen; i++) {
        Metadata.Read(x.Metadata[i], __MetadataSlice[i], reader);
    }
    x.Metadata.length = __MetadataLen;
    }

}

export function NewInvocation(): Invocation {
    let x: Invocation = {
    Source: NewSource(),
    Destination: NewDestination(),
    Payload: new Array<u8>(0),
    Metadata: new Array<Metadata>(0),
    }
    return x;
}

@unmanaged
export class MetadataViewer {
    private _0: u64;
    private _1: u64;
    private _2: u64;
    private _3: u64;

    @inline
    SizeOf(): u32 {
        return 32;
    }
    @inline
    Key(reader: karmem.Reader): string {
        let offset: u32 = load<u32>(changetype<usize>(this) + 0);
        let size: u32 = load<u32>(changetype<usize>(this) + 0 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
    @inline
    Value(reader: karmem.Reader): string {
        let offset: u32 = load<u32>(changetype<usize>(this) + 12);
        let size: u32 = load<u32>(changetype<usize>(this) + 12 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
}

@inline export function NewMetadataViewer(reader: karmem.Reader, offset: u32): MetadataViewer {
    if (!reader.IsValidOffset(offset, 32)) {
        return changetype<MetadataViewer>(changetype<usize>(_Null))
    }

    let v: MetadataViewer = changetype<MetadataViewer>(reader.Pointer + offset)
    return v
}
@unmanaged
export class SourceViewer {
    private _0: u64;
    private _1: u64;
    private _2: u64;
    private _3: u64;

    @inline
    SizeOf(): u32 {
        return load<u32>(changetype<usize>(this));
    }
    @inline
    Name(reader: karmem.Reader): string {
        if ((<u32>4 + <u32>12) > this.SizeOf()) {
            return ""
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 4);
        let size: u32 = load<u32>(changetype<usize>(this) + 4 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
    @inline
    Location(reader: karmem.Reader): string {
        if ((<u32>16 + <u32>12) > this.SizeOf()) {
            return ""
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 16);
        let size: u32 = load<u32>(changetype<usize>(this) + 16 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
}

@inline export function NewSourceViewer(reader: karmem.Reader, offset: u32): SourceViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return changetype<SourceViewer>(changetype<usize>(_Null))
    }

    let v: SourceViewer = changetype<SourceViewer>(reader.Pointer + offset)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return changetype<SourceViewer>(changetype<usize>(_Null))
    }
    return v
}
@unmanaged
export class DestinationViewer {
    private _0: u64;
    private _1: u64;
    private _2: u64;
    private _3: u64;

    @inline
    SizeOf(): u32 {
        return load<u32>(changetype<usize>(this));
    }
    @inline
    Name(reader: karmem.Reader): string {
        if ((<u32>4 + <u32>12) > this.SizeOf()) {
            return ""
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 4);
        let size: u32 = load<u32>(changetype<usize>(this) + 4 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
    @inline
    Location(reader: karmem.Reader): string {
        if ((<u32>16 + <u32>12) > this.SizeOf()) {
            return ""
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 16);
        let size: u32 = load<u32>(changetype<usize>(this) + 16 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return "";
        }
        return String.UTF8.decodeUnsafe(reader.Pointer + offset, size, false);
    }
}

@inline export function NewDestinationViewer(reader: karmem.Reader, offset: u32): DestinationViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return changetype<DestinationViewer>(changetype<usize>(_Null))
    }

    let v: DestinationViewer = changetype<DestinationViewer>(reader.Pointer + offset)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return changetype<DestinationViewer>(changetype<usize>(_Null))
    }
    return v
}
@unmanaged
export class InvocationViewer {
    private _0: u64;
    private _1: u64;
    private _2: u64;
    private _3: u64;
    private _4: u64;

    @inline
    SizeOf(): u32 {
        return load<u32>(changetype<usize>(this));
    }
    @inline
    Source(reader: karmem.Reader): SourceViewer {
        if ((<u32>4 + <u32>4) > this.SizeOf()) {
            return changetype<SourceViewer>(changetype<usize>(_Null));
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 4);
        return NewSourceViewer(reader, offset)
    }
    @inline
    Destination(reader: karmem.Reader): DestinationViewer {
        if ((<u32>8 + <u32>4) > this.SizeOf()) {
            return changetype<DestinationViewer>(changetype<usize>(_Null));
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 8);
        return NewDestinationViewer(reader, offset)
    }
    @inline
    Payload(reader: karmem.Reader): karmem.Slice<u8> {
        if ((<u32>12 + <u32>12) > this.SizeOf()) {
            return new karmem.Slice<u8>(0,0,0)
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 12);
        let size: u32 = load<u32>(changetype<usize>(this) + 12 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return new karmem.Slice<u8>(0, 0, 0);
        }
        let length = size / 1;
        return new karmem.Slice<u8>(reader.Pointer + offset, length, 1);
    }
    @inline
    Metadata(reader: karmem.Reader): karmem.Slice<MetadataViewer> {
        if ((<u32>24 + <u32>12) > this.SizeOf()) {
            return new karmem.Slice<MetadataViewer>(0,0,0)
        }
        let offset: u32 = load<u32>(changetype<usize>(this) + 24);
        let size: u32 = load<u32>(changetype<usize>(this) + 24 +4);
        if (!reader.IsValidOffset(offset, size)) {
            return new karmem.Slice<MetadataViewer>(0, 0, 0);
        }
        let length = size / 32;
        return new karmem.Slice<MetadataViewer>(reader.Pointer + offset, length, 32);
    }
}

@inline export function NewInvocationViewer(reader: karmem.Reader, offset: u32): InvocationViewer {
    if (!reader.IsValidOffset(offset, 8)) {
        return changetype<InvocationViewer>(changetype<usize>(_Null))
    }

    let v: InvocationViewer = changetype<InvocationViewer>(reader.Pointer + offset)
    if (!reader.IsValidOffset(offset, v.SizeOf())) {
        return changetype<InvocationViewer>(changetype<usize>(_Null))
    }
    return v
}
