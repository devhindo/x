


export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    const state = searchParams.get('state')
    const code = searchParams.get('code')

    if (!state || !code) {
        return new Response("Missing state or code", { status: 400 })
    }
    
    const res = await fetch(`https://api.twitter.com/2/oauth2/token`, {
        headers: {
          'Content-Type': 'application/json',
          'code': code,
          'grant_type': 'authorization_code',
          'client_id' : 'id',
          'redirect_uri' : 'redirect',
          'code_verifier' : 'verifier'
        },
      })

    
    
}

