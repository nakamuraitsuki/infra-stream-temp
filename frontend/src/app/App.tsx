import { RouterProvider } from "react-router/dom"
import { router } from "./router"
import { AppProviders } from "./providers"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { SpinnerOverlay } from "@/ui/SpinnerOverlay";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
    }
  }
});

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <SpinnerOverlay />
      <AppProviders>
        <RouterProvider router={router} />
      </AppProviders>
    </QueryClientProvider>
  )
}
