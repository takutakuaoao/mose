import {useState} from "react";

export default function PlayButton() {
    const [isPlaying, setIsPlaying] = useState(false);

    const onClick = () => setIsPlaying(!isPlaying)

    const displayMessage = {
        playing: "停止",
        stop: "再生"
    }

    return <button onClick={onClick}>
        {isPlaying ? displayMessage.playing : displayMessage.stop}
    </button>
}