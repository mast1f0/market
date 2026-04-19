/** Matches `domain.Product` JSON from the Go API. */
export type Product = {
  id: number;
  owner_id?: number;
  name: string;
  description: string;
  price: number;
  category_id?: number;
  image_url: string;
  stock?: number;
  created_at?: string;
};

/** Matches `domain.Category` JSON; `description` exists only on client fallbacks. */
export type Category = {
  id: number;
  name: string;
  created_at?: string;
  description?: string;
};
