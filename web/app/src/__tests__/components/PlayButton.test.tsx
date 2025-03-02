import {expect, test} from "vitest";
import {fireEvent, render} from "@testing-library/react";
import PlayButton from "@/components/PlayButton";

test("再生ボタンをクリックすると停止ボタンに切り替わる", () => {
    const {getByText} = render(<PlayButton/>);

    fireEvent.click(getByText("再生"));

    expect(getByText('停止')).toBeDefined();
})