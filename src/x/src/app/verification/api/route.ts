import { NextResponse } from 'next/server'

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    const state = searchParams.get('state')
    const code = searchParams.get('code')
    console.log('state', state)
    console.log('code', code)
    // return response with state and code
    return new Response(JSON.stringify({ state, code }), {
        headers: { 'content-type': 'application/json' },
    })
}
