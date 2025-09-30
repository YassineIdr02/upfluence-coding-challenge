import { useState, useEffect } from "react";
import type { AnalysisResult } from "../types/Analysisresult";
import { motion, AnimatePresence } from "motion/react"

const dimensions = ["likes", "comments", "favorites", "retweets"];
const units = ["s", "m", "h"];

export default function AnalysisForm() {
  useEffect(() => {
    document.title = "Upfluence SSE Analysis";
  }, []);
  const [durationNumber, setDurationNumber] = useState(5);
  const [durationUnit, setDurationUnit] = useState("s");
  const [dimension, setDimension] = useState(dimensions[0]);
  const [result, setResult] = useState<AnalysisResult | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchAnalysis = async () => {
    setLoading(true);
    setError(null);
    setResult(null);
    const duration = `${durationNumber}${durationUnit}`;
    try {
      const res = await fetch(
        `http://localhost:8080/analysis?duration=${duration}&dimension=${dimension}`
      );
      if (!res.ok) throw new Error(`HTTP error: ${res.status}`);
      const data: AnalysisResult = await res.json();
      setResult(data);
    } catch (err: any) {
      setError(err.message);
      setResult(null);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full flex flex-col items-center justify-center p-4">
      <h2 className="text-3xl font-bold mb-4 mt-4">Upfluence Coding Challenge - SSE Analysis</h2>
      <div className="flex flex-row w-full items-center justify-center gap-2 mb-4 font-mono font-semibold">
        /analysis?duration=
        <input
          type="number"
          min={1}
          max={100}
          value={durationNumber}
          onChange={(e) => setDurationNumber(Number(e.target.value))}
          className="input input-ghost w-fit bg-gray-200 cursor-text hover:shadow-md hover:border-gray-300 duration-700"
        />
        <select
          value={durationUnit}
          onChange={(e) => setDurationUnit(e.target.value)}
          className="select select-ghost p-1 w-[5%] bg-gray-200 cursor-pointer hover:shadow-md hover:border-gray-300 duration-700"
        >
          {units.map((u) => (
            <option key={u} value={u}>
              {u}
            </option>
          ))}
        </select>
        &dimension=
        <select
          value={dimension}
          onChange={(e) => setDimension(e.target.value)}
          className="select select-ghost p-1 w-[7%] bg-gray-200 cursor-pointer hover:shadow-md hover:border-gray-300 duration-700"
        >
          {dimensions.map((d) => (
            <option key={d} value={d}>
              {d}
            </option>
          ))}
        </select>

        <button
          onClick={fetchAnalysis}
          className="btn btn-neutral"
        >
          Analyze
        </button>
      </div>

      {loading && <> <div className="flex flex-col items-center justify-center gap-2 text-center"><p>Please wait while we analyze the data...</p><span className="loading loading-spinner loading-md"></span></div>
      </>
      }
      {error && <p className="text-red-500">{error}</p>}
      <div className="relative w-full flex items-start justify-center h-64">
        <AnimatePresence>
          {result && (
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: 20 }}
              transition={{ duration: 0.25, ease: "easeOut" }}
              className="card w-96 bg-base-100 card-md shadow-sm absolute"
            >
              <div className="card-body">
                <h2 className="card-title">Analysis Result</h2>
                <p>Total Posts: {result.total_posts}</p>
                <p>Min Timestamp: {result.minimum_timestamp}</p>
                <p>Max Timestamp: {result.maximum_timestamp}</p>
                <p>Average {dimension}: {result[`avg_${dimension}` as keyof AnalysisResult]}</p>
              </div>
            </motion.div>
          )}
        </AnimatePresence>
      </div>
    </div>
  );
}