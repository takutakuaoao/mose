import type {Meta, StoryObj} from '@storybook/react';
import AudioRecorder from "@/components/Recorder/Recorder";
import {useState} from "react";

const RecorderSandbox = () => {
    const [isPlaying, setIsPlaying] = useState(false);

    return <>
        <button onClick={() => setIsPlaying(!isPlaying)}>{isPlaying ? '録音を停止する' : '録音を始める'}</button>
        <AudioRecorder isRecording={isPlaying} />
    </>
}

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta = {
    title: 'Components/Recorder',
    component: RecorderSandbox,
    tags: ['autodocs'],
} satisfies Meta<typeof RecorderSandbox>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};
