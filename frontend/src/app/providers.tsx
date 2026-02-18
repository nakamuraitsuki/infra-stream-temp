import { type ReactNode } from "react"
import { ServiceProvider } from "../context/ServiceContext"

type Props = {
  children: ReactNode
}

export function AppProviders({ children }: Props) {
  return (
    <ServiceProvider>
      {children}
    </ServiceProvider>
  )
}
