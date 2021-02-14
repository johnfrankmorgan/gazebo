#[derive(Debug, Clone, Copy, PartialEq)]
pub enum Type {
    Invalid,
    ParenOpen,
    ParenClose,
    Comment,
    Whitespace,
    Number,
    String,
    Identifier,
}

#[derive(Debug)]
pub struct Token {
    typ: Type,
    val: String,
}

impl Token {
    pub fn new(typ: Type, val: String) -> Token {
        Token { typ: typ, val: val }
    }

    pub fn get_type(&self) -> Type {
        self.typ
    }

    pub fn get_value(&self) -> &String {
        &self.val
    }

    pub fn is(&self, typ: Type) -> bool {
        return self.get_type() == typ;
    }

    pub fn atom(&self) -> bool {
        self.is(Type::Number) || self.is(Type::String) || self.is(Type::Identifier)
    }
}
