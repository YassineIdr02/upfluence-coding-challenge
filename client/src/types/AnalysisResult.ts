export interface AnalysisResult {
    total_posts: number;
    minimum_timestamp: number;
    maximum_timestamp: number;
    [key: string]: number; 
  }