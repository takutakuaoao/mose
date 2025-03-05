export interface MediaInputStream {
    start(): void;
    stop(): void;
    onStopEvent(callBack: (e: BlobEvent) => void): void;
}

export class AudioInputStream implements MediaInputStream {
    public static async init(): Promise<AudioInputStream> {
        const stream = await navigator.mediaDevices.getUserMedia({
            audio: true,
        });

        const mediaRecorder = new MediaRecorder(stream)

        return new AudioInputStream(mediaRecorder)
    }

    constructor(public media: MediaRecorder) {
    }

    public start(): void {
        this.media.start();
    }

    public stop(): void {
        this.media.stop();
    }

    public onStopEvent(callBack: (e: BlobEvent) => void): void {
        this.media.ondataavailable = callBack;
    }
}
