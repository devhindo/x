import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'X',
  description: 'Tweet from terminal and stuff',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
