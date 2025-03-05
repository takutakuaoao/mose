import {expect, test, vi} from "vitest";
import {MediaInputStream} from "@/components/Recorder/Media";
import {render} from "@testing-library/react";
import {Recorder} from "@/components/Recorder/Recorder";

test("録音開始になるとstartメソッドが呼ばれる", () => {
    const spy =vi.spyOn(mockInput, 'start')

    render(<Recorder isRecording={true} mediaInputStream={mockInput} />)

    expect(spy).toHaveBeenCalled()
})

test("録音停止状態になるとstopメソッドが呼ばれる", () => {
    const spy = vi.spyOn(mockInput, 'stop')

    render(<Recorder isRecording={false} mediaInputStream={mockInput} />)

    expect(spy).toHaveBeenCalled()
})

const mockInput: MediaInputStream = {
    start: () => {},
    stop: () => {},
    onStopEvent: (_: (e: BlobEvent) => void)=> {
    }
}
