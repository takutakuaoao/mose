import type { Meta, StoryObj } from '@storybook/react';
import PlayButton from "@/components/PlayButton";

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta = {
    title: 'Components/PlayButton',
    component: PlayButton,
    tags: ['autodocs'],
} satisfies Meta<typeof PlayButton>;

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {};