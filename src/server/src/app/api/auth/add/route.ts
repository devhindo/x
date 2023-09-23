import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

type User = {
    state?: string
    code_verifier?: string
    code_challenge?: string
    license?: string
    auth_url?: string
}

export async function POST(request: Request) {
    const data = await request.json()
    //console.log(data)

    const user: User = data

    //const { state, code_verifier, code_challenge } = user

    return await add_user_to_supabase(user)

}

async function add_user_to_supabase(user: User) {
        const { error } = await supabase
        .from('users')
        .insert({ code_verifier: user.code_verifier, code_challenge: user.code_challenge, license: user.license, state: user.state, auth_url: user.auth_url })

        if (error) {
            return NextResponse.json({ error, message: 'err' }, { status: 500 })
        }
        return NextResponse.json({ message: 'added'}, { status: 200 })
}

export async function GET(request: Request) {
    return NextResponse.json({ message: 'Hello from the API' })
}