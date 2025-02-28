import Link from "next/link";

export default function Home() {
    // const data = await fetch(`${process.env.API_BASE_URL}/health-checker`)
    // const posts = await data.json()

    return (
        <div>
            <h1>Home</h1>
            <Link href="/about">About</Link>
        </div>
    );
}
