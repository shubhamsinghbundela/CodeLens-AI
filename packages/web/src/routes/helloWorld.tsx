import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/helloWorld")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/helloWorld"!</div>;
}
