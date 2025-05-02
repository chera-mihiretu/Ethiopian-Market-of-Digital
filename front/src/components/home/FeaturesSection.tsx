"use client";

export default function FeaturesSection() {
  return (
    <section className="py-20 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <h2 className="text-3xl sm:text-4xl font-bold text-center text-gray-900 dark:text-white mb-16">
          Why Choose Our Platform?
        </h2>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {/* Trust & Transparency */}
          <div className="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-lg transform transition-all duration-300 hover:scale-105">
            <div className="w-12 h-12 bg-blue-100 dark:bg-blue-900 rounded-full flex items-center justify-center mb-4">
              <svg className="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
            </div>
            <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-2">Trust & Transparency</h3>
            <p className="text-gray-600 dark:text-gray-300">Verified information about creators and companies in Ethiopia's digital space.</p>
          </div>

          {/* Community Power */}
          <div className="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-lg transform transition-all duration-300 hover:scale-105">
            <div className="w-12 h-12 bg-purple-100 dark:bg-purple-900 rounded-full flex items-center justify-center mb-4">
              <svg className="w-6 h-6 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-2">Community Power</h3>
            <p className="text-gray-600 dark:text-gray-300">Join a vibrant community of digital creators and innovators.</p>
          </div>

          {/* Rich Content */}
          <div className="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-lg transform transition-all duration-300 hover:scale-105">
            <div className="w-12 h-12 bg-green-100 dark:bg-green-900 rounded-full flex items-center justify-center mb-4">
              <svg className="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z" />
              </svg>
            </div>
            <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-2">Rich Content</h3>
            <p className="text-gray-600 dark:text-gray-300">Access blogs, reviews, and premium digital content.</p>
          </div>

          {/* Events & Activities */}
          <div className="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-lg transform transition-all duration-300 hover:scale-105">
            <div className="w-12 h-12 bg-red-100 dark:bg-red-900 rounded-full flex items-center justify-center mb-4">
              <svg className="w-6 h-6 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
            <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-2">Events & Activities</h3>
            <p className="text-gray-600 dark:text-gray-300">Stay updated with the latest digital events and meetups.</p>
          </div>
        </div>
      </div>
    </section>
  );
} 