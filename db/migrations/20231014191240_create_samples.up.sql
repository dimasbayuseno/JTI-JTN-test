-- Create a table to store sample data
CREATE TABLE IF NOT EXISTS sample_datas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Create an index on the 'id' column
CREATE INDEX IF NOT EXISTS idx_sample_datas_id ON sample_datas (id);

-- Create an index on the 'created_at' column
CREATE INDEX IF NOT EXISTS idx_sample_datas_created_at ON sample_datas (created_at);

-- Create an index on the 'updated_at' column
CREATE INDEX IF NOT EXISTS idx_sample_datas_updated_at ON sample_datas (updated_at);

-- Create an index on the 'deleted_at' column
CREATE INDEX IF NOT EXISTS idx_sample_datas_deleted_at ON sample_datas (deleted_at);