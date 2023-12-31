import './globals.css'
import { Inter } from 'next/font/google'
import { Header } from './common/header';
import Head from './head';

const inter = Inter({ subsets: ['latin'] })

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <Head />
      <body className={inter.className}>
        <Header />
        {children}
      </body>
    </html>
  )
}
