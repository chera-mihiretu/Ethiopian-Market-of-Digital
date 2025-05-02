"use client";

import ThemeSwitcher from "@/components/home/ThemeSwitcher";
import HeroSection from "@/components/home/HeroSection";
import StatsSection from "@/components/home/StatsSection";
import FeaturesSection from "@/components/home/FeaturesSection";
import CTASection from "@/components/home/CTASection";

export default function Home() {
  return (
    <div className="min-h-screen bg-white dark:bg-black transition-colors duration-300">
      <ThemeSwitcher />
      <HeroSection />
      <StatsSection />
      <FeaturesSection />
      <CTASection />
    </div>
  );
}
