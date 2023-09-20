import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

type User = {
    state?: string
    code_verifier?: string
    code_challenge?: string
}

export async function POST(request: Request) {
    const data = await request.json()
    console.log(data)

    const user: User = data

    const { state, code_verifier, code_challenge } = user

    return await add_user_to_supabase(user)

}

async function add_user_to_supabase(user: User) {
    // check if user exists by state

    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('state', user.state)
    .maybeSingle()

    if (error) {
        return NextResponse.json({ error })
    }

    if (data) {
        const { error } = await supabase
        .from('users')
        .update({ code_verifier: user.code_verifier, code_challenge: user.code_challenge })
        .eq('state', user.state)
    } else {
        const { data, error } = await supabase
        .from('users')
        .insert([{ state: user.state, code_verifier: user.code_verifier, code_challenge: user.code_challenge }])
    }
    return NextResponse.json({ data, error })
}