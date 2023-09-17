import { NextResponse } from 'next/server'

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    const state = searchParams.get('state')
    const code = searchParams.get('code')
    console.log('state', state)
    console.log('code', code)

    if (!state || !code) {
        return new Response("Missing state or code", { status: 400 })
    }
    // 
    return new Response(JSON.stringify({ state, code }), {
        headers: { 'content-type': 'application/json' },
    })
}

