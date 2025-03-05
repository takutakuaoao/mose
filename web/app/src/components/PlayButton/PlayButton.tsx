import {useState} from "react";

type Props = {
    onClick: (isPlaying: boolean) => void;
}

export default function PlayButton(props: Props) {
    const [isPlaying, setIsPlaying] = useState(false);

    const onClick = () => {
        props.onClick(!isPlaying);
        setIsPlaying(!isPlaying);
    }

    const displayMessage = {
        playing: "停止",
        stop: "再生"
    }

    return <button onClick={onClick}>
        {isPlaying ? displayMessage.playing : displayMessage.stop}
    </button>
}