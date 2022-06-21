namespace Models {
    public abstract class Model {
        public object Id {get; set; }

        protected Model(object id)
        {
            Id = id;
        }
    }
}