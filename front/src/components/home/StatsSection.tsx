"use client";

export default function StatsSection() {
  return (
    <section className="py-20 px-4 sm:px-6 lg:px-8 bg-gray-50 dark:bg-gray-900">
      <div className="max-w-7xl mx-auto">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div className="text-center">
            <div className="text-4xl font-bold text-blue-600 dark:text-blue-400 mb-2">500+</div>
            <div className="text-gray-600 dark:text-gray-300">Digital Creators</div>
          </div>
          <div className="text-center">
            <div className="text-4xl font-bold text-purple-600 dark:text-purple-400 mb-2">200+</div>
            <div className="text-gray-600 dark:text-gray-300">Companies</div>
          </div>
          <div className="text-center">
            <div className="text-4xl font-bold text-green-600 dark:text-green-400 mb-2">1000+</div>
            <div className="text-gray-600 dark:text-gray-300">Digital Products</div>
          </div>
          <div className="text-center">
            <div className="text-4xl font-bold text-red-600 dark:text-red-400 mb-2">50+</div>
            <div className="text-gray-600 dark:text-gray-300">Events Monthly</div>
          </div>
        </div>
      </div>
    </section>
  );
} 