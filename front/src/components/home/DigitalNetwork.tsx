"use client";

export default function DigitalNetwork() {
  // Fixed positions for nodes
  const nodes = [
    { cx: 20, cy: 30, delay: 0.2, duration: 2.5 },
    { cx: 40, cy: 60, delay: 0.4, duration: 2.8 },
    { cx: 60, cy: 20, delay: 0.6, duration: 3.0 },
    { cx: 80, cy: 50, delay: 0.8, duration: 2.7 },
    { cx: 30, cy: 80, delay: 1.0, duration: 2.9 },
    { cx: 70, cy: 70, delay: 1.2, duration: 2.6 },
    { cx: 10, cy: 40, delay: 1.4, duration: 2.4 },
    { cx: 50, cy: 10, delay: 1.6, duration: 2.3 },
    { cx: 90, cy: 90, delay: 1.8, duration: 2.2 },
    { cx: 25, cy: 55, delay: 2.0, duration: 2.1 },
  ];

  // Fixed positions for connections
  const connections = [
    { x1: 20, y1: 30, x2: 40, y2: 60, delay: 0.3, duration: 3.0 },
    { x1: 40, y1: 60, x2: 60, y2: 20, delay: 0.5, duration: 2.8 },
    { x1: 60, y1: 20, x2: 80, y2: 50, delay: 0.7, duration: 2.9 },
    { x1: 80, y1: 50, x2: 30, y2: 80, delay: 0.9, duration: 2.7 },
    { x1: 30, y1: 80, x2: 70, y2: 70, delay: 1.1, duration: 2.6 },
    { x1: 70, y1: 70, x2: 10, y2: 40, delay: 1.3, duration: 2.5 },
    { x1: 10, y1: 40, x2: 50, y2: 10, delay: 1.5, duration: 2.4 },
    { x1: 50, y1: 10, x2: 90, y2: 90, delay: 1.7, duration: 2.3 },
  ];

  // Fixed positions for floating elements
  const floatingElements = [
    { left: 15, top: 25, delay: 0.2, duration: 12 },
    { left: 45, top: 65, delay: 0.4, duration: 14 },
    { left: 75, top: 35, delay: 0.6, duration: 13 },
    { left: 25, top: 85, delay: 0.8, duration: 15 },
    { left: 85, top: 15, delay: 1.0, duration: 11 },
    { left: 55, top: 45, delay: 1.2, duration: 16 },
  ];

  return (
    <div className="absolute inset-0 overflow-hidden">
      <svg className="absolute w-full h-full" viewBox="0 0 100 100" preserveAspectRatio="none">
        {/* Animated Grid Lines */}
        <pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse">
          <path d="M 10 0 L 0 0 0 10" fill="none" stroke="currentColor" strokeWidth="0.5" className="text-blue-500/20 dark:text-blue-400/20" />
        </pattern>
        <rect width="100" height="100" fill="url(#grid)" />
        
        {/* Animated Nodes */}
        {nodes.map((node, i) => (
          <circle
            key={i}
            cx={node.cx}
            cy={node.cy}
            r="0.5"
            className="text-blue-500 dark:text-blue-400"
            style={{
              animation: `pulse ${node.duration}s infinite ${node.delay}s`,
            }}
          />
        ))}

        {/* Animated Connections */}
        {connections.map((conn, i) => (
          <line
            key={i}
            x1={conn.x1}
            y1={conn.y1}
            x2={conn.x2}
            y2={conn.y2}
            stroke="currentColor"
            strokeWidth="0.2"
            className="text-blue-500/30 dark:text-blue-400/30"
            style={{
              animation: `fade ${conn.duration}s infinite ${conn.delay}s`,
            }}
          />
        ))}
      </svg>

      {/* Floating Elements */}
      <div className="absolute inset-0">
        {floatingElements.map((elem, i) => (
          <div
            key={i}
            className="absolute w-4 h-4 rounded-full bg-blue-500/10 dark:bg-blue-400/10"
            style={{
              left: `${elem.left}%`,
              top: `${elem.top}%`,
              animation: `float ${elem.duration}s infinite ${elem.delay}s`,
            }}
          />
        ))}
      </div>
    </div>
  );
} 