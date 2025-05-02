"use client";

import Link from "next/link";
import DigitalNetwork from "./DigitalNetwork";

export default function HeroSection() {
  return (
    <section className="relative min-h-screen flex items-center justify-center px-4 sm:px-6 lg:px-8 overflow-hidden">
      <DigitalNetwork />
      
      {/* Gradient Overlay */}
      <div className="absolute inset-0 bg-gradient-to-br from-blue-50/50 to-purple-50/50 dark:from-gray-900/50 dark:to-black/50"></div>

      <div className="max-w-7xl mx-auto text-center relative z-10">
        <h1 className="text-5xl sm:text-7xl font-bold text-gray-900 dark:text-white mb-6 animate-fade-in">
          Explore Ethiopia's Digital Landscape
        </h1>
        <p className="text-xl sm:text-2xl text-gray-600 dark:text-gray-300 mb-12 animate-fade-in-delay">
          Discover real people, projects, and companies shaping the digital future.
        </p>
        <div className="flex flex-col sm:flex-row gap-4 justify-center items-center animate-fade-in-delay-2">
          <Link 
            href="/login"
            className="inline-block bg-blue-600 hover:bg-blue-700 text-white font-semibold px-8 py-4 rounded-full text-lg transition-all duration-300 transform hover:scale-105 shadow-lg hover:shadow-xl"
          >
            Get Started
          </Link>
          <Link 
            href="/explore"
            className="inline-block bg-white dark:bg-gray-800 text-gray-900 dark:text-white font-semibold px-8 py-4 rounded-full text-lg transition-all duration-300 transform hover:scale-105 shadow-lg hover:shadow-xl border border-gray-200 dark:border-gray-700"
          >
            Explore Platform
          </Link>
        </div>
      </div>
    </section>
  );
} 