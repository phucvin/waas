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
  let result = ping(inv.Payload(kmReader)[0]);
  if (result.isOk) {
    let resBytes = Uint8Array.wrap(new ArrayBuffer(1));
    resBytes[0] = result.get();
    return Result.ok<ArrayBuffer>(resBytes.buffer);
  } else {
    return Result.error<ArrayBuffer>(changetype<Error>(result.error()))
  }
}

function ping(countLeft: u8): Result<u8> {
  if (countLeft <= 1 ) {
    return Result.ok<u8>(0);
  }

  return invokePong(countLeft - 1);
}

function invokePong(countLeft: u8): Result<u8> {
  let inv = waaskm.NewInvocation();
  inv.Source.Name = "ping";
  inv.Source.Location = "global";
  inv.Destination.Name = "ping";
  inv.Destination.Location = "anywhere";
  inv.Payload = new Array<u8>(1);
  inv.Payload[0] = countLeft;
  kmWriter.Reset();
  let writeOk = waaskm.Invocation.Write(inv, kmWriter, 0);
	if (!writeOk) {
    return Result.error<u8>(new Error("write failed"));
	}

  let result = hostCall("", "", "invoke", kmWriter.Bytes().buffer);
  if (result.isOk) {
    return Result.ok<u8>(Uint8Array.wrap(result.get())[0]);
  } else {
    return Result.error<u8>(changetype<Error>(result.error()));
  }
}

// This must be present in the entry file. Do not remove.

export function __guest_call(operation_size: usize, payload_size: usize): bool {
  return handleCall(operation_size, payload_size);
}

function abort(message: string | null, fileName: string | null, lineNumber: u32, columnNumber: u32): void {
  handleAbort(message, fileName, lineNumber, columnNumber)
}