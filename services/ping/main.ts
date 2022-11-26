import {
  register,
  handleCall,
  hostCall,
  handleAbort,
  Result,
} from "@wapc/as-guest";
import * as karmem from "karmem/assemblyscript/karmem.ts";
import * as waaskm from "../../km/waas_generated.ts";

const kmWriter = karmem.NewWriter(1024);

register("ping", pingWrapper);

function pingWrapper(payload: ArrayBuffer): Result<ArrayBuffer> {
  let kmReader = karmem.NewReader(changetype<StaticArray<u8>>(payload));
  let inv = waaskm.NewInvocationViewer(kmReader, 0);
  let result = ping(inv.Payload(kmReader));
  if (result.isOk) {
    return Result.ok<ArrayBuffer>(result.get().buffer);
  } else {
    return Result.error<ArrayBuffer>(changetype<Error>(result.error()))
  }
}

function ping(counts: karmem.Slice<u8>): Result<Uint8Array> {
  if (counts[0] == 0) {
    let sameCounts = new Uint8Array(counts.length);
    for (let i = 0; i < counts.length; i++) {
      sameCounts[i] = counts[i];
    }
    return Result.ok(sameCounts);
  }

  let newCounts = new Uint8Array(counts.length + 1);
  newCounts[0] = counts[0] == 0 ? counts[0] : (counts[0] - 1);
  for (let i = 0; i < counts.length; i++) {
    newCounts[i + 1] = counts[i];
  }
  return invokePong(newCounts);
}

function invokePong(counts: Uint8Array): Result<Uint8Array> {
  let inv = waaskm.NewInvocation();
  inv.Source.Name = "ping";
  inv.Source.Location = "global";
  inv.Destination.Name = "ping";
  inv.Destination.Location = "anywhere";
  inv.Payload = new Array<u8>(counts.length);
  for (let i = 0; i < counts.length; ++i) {
    inv.Payload[i] = counts[i];
  }
  kmWriter.Reset();
  let writeOk = waaskm.Invocation.Write(inv, kmWriter, 0);
	if (!writeOk) {
    return Result.error<Uint8Array>(new Error("write failed"));
	}

  let result = hostCall("", "", "invoke", kmWriter.Bytes().buffer);
  if (result.isOk) {
    return Result.ok(Uint8Array.wrap(result.get()))
  } else {
    return Result.error<Uint8Array>(changetype<Error>(result.error()));
  }
}

// This must be present in the entry file. Do not remove.

export function __guest_call(operation_size: usize, payload_size: usize): bool {
  return handleCall(operation_size, payload_size);
}

function abort(message: string | null, fileName: string | null, lineNumber: u32, columnNumber: u32): void {
  handleAbort(message, fileName, lineNumber, columnNumber)
}