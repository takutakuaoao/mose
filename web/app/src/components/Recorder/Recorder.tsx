import {useEffect, useRef, useState} from "react";
import {AudioInputStream, MediaInputStream} from "@/components/Recorder/Media";

type Props = {
    isRecording: boolean;
    mediaInputStream?: MediaInputStream
}

const AudioRecorder = (props: Omit<Props, 'mediaInputStream'>) => {
    const [mediaInputStream, setMediaInputStream] = useState<MediaInputStream | undefined>(undefined);

    useEffect(() => {
        const init = async () => {
            const input = await AudioInputStream.init()

            setMediaInputStream(input)
        }

        init()
    }, []);

    return <Recorder {...props} mediaInputStream={mediaInputStream} />;
}

export default AudioRecorder;

export function Recorder(props: Props) {
    const [srcURL, setSrcURL] = useState("");
    const inputStreamRef = useRef<MediaInputStream>(null);

    useEffect(() => {
        const initInputStream = (mediaInputStream?: MediaInputStream): MediaInputStream | null => {
            if (!isDefinedInputStream(mediaInputStream)) {
                return null;
            }

            mediaInputStream.onStopEvent((e) => {
                setSrcURL(URL.createObjectURL(e.data));
            })

            return mediaInputStream;
        }

        inputStreamRef.current = initInputStream(props.mediaInputStream);
    }, [props.mediaInputStream]);

    useEffect(() => {
        if (!isDefinedInputStream(inputStreamRef.current)) {
            return;
        }

        if (props.isRecording) {
            inputStreamRef.current.start()
        } else {
            inputStreamRef.current.stop()
        }
    }, [props.isRecording]);

    function isDefinedInputStream(mediaInputStream?: MediaInputStream | null): mediaInputStream is MediaInputStream {
        return !(mediaInputStream === undefined || mediaInputStream === null);
    }

    return <>
        <audio src={srcURL} controls={true}/>
    </>
}
