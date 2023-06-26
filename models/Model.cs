namespace Models {
    public abstract class Model<T> {
        public T Id {get; set; }

        protected Model(object id)
        {
            Id = id;
        }

        public Model() {}
    }
}