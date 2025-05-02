"use client";

import Link from "next/link";

export default function CTASection() {
  return (
    <section className="py-20 px-4 sm:px-6 lg:px-8 bg-blue-600 dark:bg-blue-900">
      <div className="max-w-7xl mx-auto text-center">
        <h2 className="text-3xl sm:text-4xl font-bold text-white mb-8">
          Ready to Join Ethiopia's Digital Revolution?
        </h2>
        <Link 
          href="/signup"
          className="inline-block bg-white text-blue-600 hover:bg-gray-100 font-semibold px-8 py-4 rounded-full text-lg transition-all duration-300 transform hover:scale-105 shadow-lg hover:shadow-xl"
        >
          Create Your Profile
        </Link>
      </div>
    </section>
  );
} 