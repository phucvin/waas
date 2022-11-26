import {
  register,
  handleCall,
  hostCall,
  handleAbort,
  Result,
} from "@wapc/as-guest";
import * as karmem from "karmem/assemblyscript/karmem.ts";
import * as waaskm from "../../km/waas_generated.ts";

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
  return Result.ok(countLeft - 1);
}

// This must be present in the entry file. Do not remove.

export function __guest_call(operation_size: usize, payload_size: usize): bool {
  return handleCall(operation_size, payload_size);
}

function abort(message: string | null, fileName: string | null, lineNumber: u32, columnNumber: u32): void {
  handleAbort(message, fileName, lineNumber, columnNumber)
}