import { NextResponse } from "next/server"

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    
    const state = searchParams.get("state")
    const code = searchParams.get("code")

    if (!state || !code) {
        return new NextResponse("Missing state or code | Authentication failed", { status: 400 })
    }
    return NextResponse.json({ state, code })
}

type User = {
    username?: string
    secret?:string
}

export async function POST(request: Request) {
    const data: User = await request.json()
    console.log(data)

    const { username, secret } = data

    return NextResponse.json({ username, secret })
}
