import {expect, test} from "vitest";
import {fireEvent, render} from "@testing-library/react";
import PlayButton from "@/components/PlayButton/PlayButton";

test("再生ボタンをクリックすると停止ボタンに切り替わる", () => {
    const {getByText} = render(<PlayButton onClick={() => {}} />);

    fireEvent.click(getByText("再生"));

    expect(getByText('停止')).toBeDefined();
})

test("再生ボタンをクリックすると関数が呼ばれる", () => {
    let isCalled = false;
    const play = (isPlaying: boolean) => {
        if (isPlaying) {
            isCalled = true;
        }
    }

    const {getByText} = render(<PlayButton onClick={play} />);

    fireEvent.click(getByText("再生"));

    expect(isCalled).toBe(true);
})