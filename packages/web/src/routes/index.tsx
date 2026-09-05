import Container from "#/components/common/container";
import { ModeToggle } from "#/components/ui/mode-toggle";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <nav className="sticky top-4 z-50 mb-6 px-4">
        <Container className="container mx-auto">
          <div className="flex items-center justify-between rounded-[12px] border border-[#333234]/15 bg-white/95 px-6 py-3 shadow-2xs backdrop-blur-md">
            <div className="flex items-center gap-2.5">
              <div className="flex h-7 w-7 items-center justify-center rounded-[8px] bg-[#333234] font-display text-xs font-bold text-[#ebeadf]">
                <ModeToggle />
              </div>
              <h1 className="font-display text-xl font-bold tracking-tight text-[#333234]">
                swags.me
              </h1>
            </div>
          </div>
        </Container>
      </nav>
    </div>
  );
}
